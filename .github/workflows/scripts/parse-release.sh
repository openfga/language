#!/usr/bin/env bash
# Extracted from .github/workflows/reusable-release-please.yaml so the parsing
# logic can be unit-tested independently of GitHub Actions.
#
# Usage:
#   parse-release.sh manifest-diff <current.json> <previous.json>
#     -> prints JSON array of {component, version, tag_name} for changed pkgs
#        exits 1 if no changes detected
#
#   parse-release.sh changelog-notes <changelog-path> <version>
#     -> prints the release notes section for the given version
#        falls back to "Release <version>" if no section found
#
#   parse-release.sh next-version <base-version> <patch|minor|major>
#     -> prints the next semantic version after <base-version> for the given
#        bump type (pre-release/build metadata on <base-version> is dropped).
#        Used by the workflow_dispatch path to force a Release-As version.
set -euo pipefail

cmd="${1:-}"
shift || true

case "$cmd" in
  manifest-diff)
    current_file="$1"
    previous_file="$2"

    releases=$(jq -c -n \
      --argjson cur  "$(cat "$current_file")" \
      --argjson prev "$(cat "$previous_file")" \
      '[ $cur | to_entries[]
         | select(.value != $prev[.key])
         | { component: .key, version: .value, tag_name: "\(.key)/v\(.value)" } ]')

    if [[ $(jq 'length' <<<"$releases") -eq 0 ]]; then
      echo "::error::No version changes detected." >&2
      exit 1
    fi

    echo "$releases"
    ;;

  changelog-notes)
    changelog="$1"
    version="$2"

    notes=$(awk -v ver="$version" '
      BEGIN {
        gsub(/\./, "\\.", ver)
        pattern = "(^|[^0-9A-Za-z.-])" ver "([^0-9A-Za-z.-]|$)"
      }
      /^## / {
        if (found) exit
        if ($0 ~ pattern) { found=1; next }
      }
      found { print }
    ' "$changelog")

    if [[ -z "$notes" ]]; then
      echo "Release ${version}"
    else
      echo "$notes"
    fi
    ;;

  next-version)
    base="$1"
    bump="$2"

    # Strip a leading "v" and any pre-release / build metadata.
    base="${base#v}"
    base="${base%%-*}"
    base="${base%%+*}"

    IFS='.' read -r major minor patch <<<"$base"

    case "$bump" in
      major) major=$((major + 1)); minor=0; patch=0 ;;
      minor) minor=$((minor + 1)); patch=0 ;;
      patch) patch=$((patch + 1)) ;;
      *) echo "Usage: $0 next-version <base-version> <patch|minor|major>" >&2; exit 2 ;;
    esac

    echo "${major}.${minor}.${patch}"
    ;;

  *)
    echo "Usage: $0 {manifest-diff|changelog-notes|next-version} ..." >&2
    exit 2
    ;;
esac

