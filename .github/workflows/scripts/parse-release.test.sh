#!/usr/bin/env bash
#
# parse-release_test.sh - lightweight tests for parse-release.sh.
#
# Usage: .github/workflows/scripts/parse-release_test.sh
# Requires: bash, jq
#
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PARSE="${SCRIPT_DIR}/parse-release.sh"

TMP="$(mktemp -d)"
trap 'rm -rf "$TMP"' EXIT

pass=0
fail=0

assert_eq() {
  local name="$1" expected="$2" actual="$3"
  if [[ "$expected" == "$actual" ]]; then
    pass=$((pass + 1))
    echo "ok   - $name"
  else
    fail=$((fail + 1))
    echo "FAIL - $name"
    echo "       expected: |$expected|"
    echo "       actual:   |$actual|"
  fi
}

##### Fixtures #####

cat >"$TMP/prev.json" <<'EOF'
{
  "pkg/go": "0.2.1",
  "pkg/js": "0.2.1",
  "pkg/java": "0.2.0-beta.2"
}
EOF

cat >"$TMP/cur.json" <<'EOF'
{
  "pkg/go": "0.2.1",
  "pkg/js": "0.2.3",
  "pkg/java": "0.3.0"
}
EOF

cat >"$TMP/CHANGELOG.md" <<'EOF'
# Changelog

## pkg/js/v0.2.3

### [v0.2.3](https://example.com/compare/pkg/js/v0.2.1...pkg/js/v0.2.3) (2026-06-04)

Added:

- A shiny new feature (#123)

Fixed:

- A pesky bug (#124)

## pkg/js/v0.2.1

### [v0.2.1](https://example.com) (2026-02-11)

Added:

- Older stuff
EOF

##### manifest-diff #####

# Only pkg/js and pkg/java changed; output sorted by component for stable compare.
diff_out="$("$PARSE" manifest-diff "$TMP/cur.json" "$TMP/prev.json" | jq -S -c 'sort_by(.component)')"
expected_diff='[{"component":"pkg/java","tag_name":"pkg/java/v0.3.0","version":"0.3.0"},{"component":"pkg/js","tag_name":"pkg/js/v0.2.3","version":"0.2.3"}]'
assert_eq "manifest-diff reports only changed components" "$expected_diff" "$diff_out"

# Missing previous manifest => everything is new.
all_new="$("$PARSE" manifest-diff "$TMP/cur.json" "$TMP/does-not-exist.json" | jq -c 'length')"
assert_eq "manifest-diff treats missing previous manifest as all-new" "3" "$all_new"

# No changes => empty array.
no_change="$("$PARSE" manifest-diff "$TMP/cur.json" "$TMP/cur.json")"
assert_eq "manifest-diff with identical manifests is empty" "[]" "$no_change"

##### changelog-notes #####

notes="$("$PARSE" changelog-notes "$TMP/CHANGELOG.md" "0.2.3")"
expected_notes='### [v0.2.3](https://example.com/compare/pkg/js/v0.2.1...pkg/js/v0.2.3) (2026-06-04)

Added:

- A shiny new feature (#123)

Fixed:

- A pesky bug (#124)'
assert_eq "changelog-notes extracts the matching section body" "$expected_notes" "$notes"

##### next-version #####

assert_eq "next-version patch" "0.2.2" "$("$PARSE" next-version 0.2.1 patch)"
assert_eq "next-version minor" "0.3.0" "$("$PARSE" next-version 0.2.1 minor)"
assert_eq "next-version major" "1.0.0" "$("$PARSE" next-version 0.2.1 major)"
assert_eq "next-version drops pre-release metadata" "0.3.0" "$("$PARSE" next-version v0.2.0-beta.2 minor)"

##### Summary #####

echo
echo "passed: $pass, failed: $fail"
[[ "$fail" -eq 0 ]]

