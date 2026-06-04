#!/usr/bin/env bash
#
# parse-release.sh - helpers used by the release-please automation.
#
# Subcommands:
#   manifest-diff <current_manifest> <previous_manifest>
#       Print a JSON array of the components whose version changed between two
#       release-please manifests. Each entry has the shape:
#         {"component":"pkg/js","version":"0.2.3","tag_name":"pkg/js/v0.2.3"}
#
#   changelog-notes <changelog_path> <version>
#       Print the markdown body of the section that documents <version>.
#       Leading/trailing blank lines are trimmed.
#
#   next-version <base_version> <patch|minor|major>
#       Print the next semantic version after <base_version> for the given bump
#       type. Any pre-release / build metadata on <base_version> is dropped.
#
set -euo pipefail

usage() {
  cat >&2 <<'EOF'
Usage:
  parse-release.sh manifest-diff   <current_manifest> <previous_manifest>
  parse-release.sh changelog-notes <changelog_path>   <version>
  parse-release.sh next-version    <base_version>     <patch|minor|major>
EOF
  exit 64
}

manifest_diff() {
  local current="$1" previous="$2"

  [[ -f "$current" ]] || { echo "current manifest not found: $current" >&2; exit 1; }
  # A missing previous manifest is treated as "everything is new".
  if [[ ! -f "$previous" ]]; then
    previous="/dev/null"
  fi

  jq -nc \
    --slurpfile cur "$current" \
    --slurpfile prev "$previous" '
      ($cur[0] // {})  as $c
    | ($prev[0] // {}) as $p
    | [ $c
        | to_entries[]
        | select(.value != ($p[.key] // null))
        | { component: .key,
            version: .value,
            tag_name: (.key + "/v" + .value) } ]
  '
}

changelog_notes() {
  local file="$1" version="$2"

  [[ -f "$file" ]] || { echo "changelog not found: $file" >&2; exit 1; }

  awk -v ver="$version" '
    function is_header(line) { return line ~ /^## / }
    BEGIN { found = 0 }
    {
      if (is_header($0)) {
        if (found) { exit }
        # Match the section header for this version (with or without a "v"
        # prefix, e.g. "## [0.2.1]" or "## pkg/go/v0.2.1").
        if (index($0, "v" ver) > 0 || index($0, ver) > 0) {
          found = 1
          next
        }
      }
      if (found) { lines[n++] = $0 }
    }
    END {
      # Trim leading blank lines.
      start = 0
      while (start < n && lines[start] ~ /^[[:space:]]*$/) start++
      # Trim trailing blank lines.
      end = n - 1
      while (end >= start && lines[end] ~ /^[[:space:]]*$/) end--
      for (i = start; i <= end; i++) print lines[i]
    }
  ' "$file"
}

next_version() {
  local base="$1" bump="$2"

  # Strip a leading "v" and any pre-release / build metadata.
  base="${base#v}"
  base="${base%%-*}"
  base="${base%%+*}"

  if [[ ! "$base" =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    echo "invalid base version: $1" >&2
    exit 1
  fi

  local major minor patch
  IFS='.' read -r major minor patch <<<"$base"

  case "$bump" in
    major) major=$((major + 1)); minor=0; patch=0 ;;
    minor) minor=$((minor + 1)); patch=0 ;;
    patch) patch=$((patch + 1)) ;;
    *) echo "invalid bump type: $bump (expected patch|minor|major)" >&2; exit 1 ;;
  esac

  printf '%s.%s.%s\n' "$major" "$minor" "$patch"
}

main() {
  [[ $# -ge 1 ]] || usage
  local cmd="$1"; shift

  case "$cmd" in
    manifest-diff)   [[ $# -eq 2 ]] || usage; manifest_diff "$@" ;;
    changelog-notes) [[ $# -eq 2 ]] || usage; changelog_notes "$@" ;;
    next-version)    [[ $# -eq 2 ]] || usage; next_version "$@" ;;
    *) usage ;;
  esac
}

main "$@"

