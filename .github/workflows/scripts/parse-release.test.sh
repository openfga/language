#!/usr/bin/env bash
# Tests for parse-release.sh
# Run: bash .github/workflows/scripts/parse-release.test.sh

set -uo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PARSE="$SCRIPT_DIR/parse-release.sh"
TMP="$(mktemp -d)"
trap 'rm -rf "$TMP"' EXIT

PASS=0
FAIL=0

assert_eq() {
  local name="$1" expected="$2" actual="$3"
  if [[ "$expected" == "$actual" ]]; then
    echo "  PASS - $name"
    PASS=$((PASS + 1))
  else
    echo "  FAIL - $name"
    echo "     expected: $expected"
    echo "     actual:   $actual"
    FAIL=$((FAIL + 1))
  fi
}

assert_contains() {
  local name="$1" needle="$2" haystack="$3"
  if [[ "$haystack" == *"$needle"* ]]; then
    echo "  PASS - $name"
    PASS=$((PASS + 1))
  else
    echo "  FAIL - $name"
    echo "     expected to contain: $needle"
    echo "     actual:              $haystack"
    FAIL=$((FAIL + 1))
  fi
}

assert_exit_code() {
  local name="$1" expected="$2" actual="$3"
  if [[ "$expected" -eq "$actual" ]]; then
    echo "  PASS - $name (exit=$actual)"
    PASS=$((PASS + 1))
  else
    echo "  FAIL - $name"
    echo "     expected exit: $expected"
    echo "     actual exit:   $actual"
    FAIL=$((FAIL + 1))
  fi
}

#####################################################################
# manifest-diff tests
#####################################################################
echo
echo "=== manifest-diff ==="

diff_run() {
  echo "$1" >"$TMP/cur.json"
  echo "$2" >"$TMP/prev.json"
  "$PARSE" manifest-diff "$TMP/cur.json" "$TMP/prev.json" 2>/dev/null
}

# 1. Only pkg/go bumped
out=$(diff_run \
  '{"pkg/go":"0.3.0","pkg/js":"0.2.4","pkg/java":"0.2.0-beta.0"}' \
  '{"pkg/go":"0.2.2","pkg/js":"0.2.4","pkg/java":"0.2.0-beta.0"}')
assert_eq "only pkg/go bumped" \
  '[{"component":"pkg/go","version":"0.3.0","tag_name":"pkg/go/v0.3.0"}]' \
  "$out"

# 2. Only pkg/js bumped
out=$(diff_run \
  '{"pkg/go":"0.2.2","pkg/js":"0.2.5","pkg/java":"0.2.0-beta.0"}' \
  '{"pkg/go":"0.2.2","pkg/js":"0.2.4","pkg/java":"0.2.0-beta.0"}')
assert_eq "only pkg/js bumped" \
  '[{"component":"pkg/js","version":"0.2.5","tag_name":"pkg/js/v0.2.5"}]' \
  "$out"

# 3. Only pkg/java bumped
out=$(diff_run \
  '{"pkg/go":"0.2.2","pkg/js":"0.2.4","pkg/java":"0.2.0-beta.1"}' \
  '{"pkg/go":"0.2.2","pkg/js":"0.2.4","pkg/java":"0.2.0-beta.0"}')
assert_eq "only pkg/java bumped (pre-release)" \
  '[{"component":"pkg/java","version":"0.2.0-beta.1","tag_name":"pkg/java/v0.2.0-beta.1"}]' \
  "$out"

# 4. Two packages bumped simultaneously
out=$(diff_run \
  '{"pkg/go":"0.3.0","pkg/js":"0.2.5","pkg/java":"0.2.0-beta.0"}' \
  '{"pkg/go":"0.2.2","pkg/js":"0.2.4","pkg/java":"0.2.0-beta.0"}')
assert_eq "pkg/go + pkg/js bumped together" \
  '[{"component":"pkg/go","version":"0.3.0","tag_name":"pkg/go/v0.3.0"},{"component":"pkg/js","version":"0.2.5","tag_name":"pkg/js/v0.2.5"}]' \
  "$out"

# 5. All three packages bumped
out=$(diff_run \
  '{"pkg/go":"0.3.0","pkg/js":"0.3.0","pkg/java":"0.3.0"}' \
  '{"pkg/go":"0.2.2","pkg/js":"0.2.4","pkg/java":"0.2.0-beta.0"}')
n=$(echo "$out" | jq 'length')
assert_eq "all three packages bumped (count=3)" "3" "$n"

# 6. No changes; expect exit 1 with error
echo '{"pkg/go":"0.2.2","pkg/js":"0.2.4","pkg/java":"0.2.0-beta.0"}' >"$TMP/cur.json"
cp "$TMP/cur.json" "$TMP/prev.json"
err=$("$PARSE" manifest-diff "$TMP/cur.json" "$TMP/prev.json" 2>&1 >/dev/null)
code=$?
assert_exit_code "no changes exits non-zero" 1 "$code"
assert_contains "no changes error message" "No version changes detected" "$err"

# 7. New package added (not in previous); expect a release
out=$(diff_run \
  '{"pkg/go":"0.2.2","pkg/js":"0.2.4","pkg/java":"0.2.0-beta.0","pkg/rust":"0.1.0"}' \
  '{"pkg/go":"0.2.2","pkg/js":"0.2.4","pkg/java":"0.2.0-beta.0"}')
assert_eq "new package added to manifest" \
  '[{"component":"pkg/rust","version":"0.1.0","tag_name":"pkg/rust/v0.1.0"}]' \
  "$out"

# 8. Package removed from manifest is ignored (not in current)
echo '{"pkg/go":"0.2.2","pkg/js":"0.2.4"}' >"$TMP/cur.json"
echo '{"pkg/go":"0.2.2","pkg/js":"0.2.4","pkg/java":"0.2.0-beta.0"}' >"$TMP/prev.json"
"$PARSE" manifest-diff "$TMP/cur.json" "$TMP/prev.json" >/dev/null 2>&1
code=$?
assert_exit_code "removed package does not trigger release" 1 "$code"

# 9. Pre-release to stable bump
out=$(diff_run \
  '{"pkg/java":"1.0.0"}' \
  '{"pkg/java":"1.0.0-rc.1"}')
assert_eq "pre-release to stable bump" \
  '[{"component":"pkg/java","version":"1.0.0","tag_name":"pkg/java/v1.0.0"}]' \
  "$out"

# 10. Patch bump that overlaps a substring of the older version
out=$(diff_run \
  '{"pkg/js":"0.2.40"}' \
  '{"pkg/js":"0.2.4"}')
assert_eq "version 0.2.4 -> 0.2.40 detected as change" \
  '[{"component":"pkg/js","version":"0.2.40","tag_name":"pkg/js/v0.2.40"}]' \
  "$out"

#####################################################################
# changelog-notes tests
#####################################################################
echo
echo "=== changelog-notes ==="

cat >"$TMP/CHANGELOG.md" <<'EOF'
# Changelog

## Unreleased

## [0.2.40](https://github.com/foo/bar/compare/pkg/js/v0.2.39...pkg/js/v0.2.40) (2026-06-01)


### Added

* feature for 0.2.40 ([#999](https://github.com/foo/bar/issues/999))


## [0.2.4](https://github.com/foo/bar/compare/pkg/js/v0.2.3...pkg/js/v0.2.4) (2026-05-28)

> [!NOTE]
> Manual note added between version headings.

### Fixed

* bug fix for 0.2.4 ([#283](https://github.com/foo/bar/issues/283))


## [0.2.3](https://github.com/foo/bar/compare/pkg/js/v0.2.2...pkg/js/v0.2.3) (2026-05-26)


### Miscellaneous

* release 0.2.3


## [0.2.0-beta.1](https://github.com/foo/bar/compare/pkg/js/v0.2.0-beta.0...pkg/js/v0.2.0-beta.1) (2026-04-10)


### Added

* beta feature
EOF

# 1. Latest version captures its own section and ### subheadings
out=$("$PARSE" changelog-notes "$TMP/CHANGELOG.md" "0.2.40")
assert_contains "0.2.40 captures ### Added subheading" "### Added" "$out"
assert_contains "0.2.40 captures its body" "feature for 0.2.40" "$out"
if [[ "$out" != *"0.2.4 "* && "$out" != *"v0.2.4)"* ]]; then
  echo "  PASS - 0.2.40 does not leak into 0.2.4 section"
  PASS=$((PASS + 1))
else
  echo "  FAIL - 0.2.40 leaked into 0.2.4 section: $out"
  FAIL=$((FAIL + 1))
fi

# 2. Substring collision: 0.2.4 must not capture 0.2.40
out=$("$PARSE" changelog-notes "$TMP/CHANGELOG.md" "0.2.4")
assert_contains "0.2.4 captures manual NOTE block" "Manual note added" "$out"
assert_contains "0.2.4 captures ### Fixed" "### Fixed" "$out"
assert_contains "0.2.4 captures bug fix line" "bug fix for 0.2.4" "$out"
if [[ "$out" == *"feature for 0.2.40"* ]]; then
  echo "  FAIL - 0.2.4 incorrectly captured 0.2.40 content"
  FAIL=$((FAIL + 1))
else
  echo "  PASS - 0.2.4 does not capture 0.2.40 content"
  PASS=$((PASS + 1))
fi

# 3. Older version
out=$("$PARSE" changelog-notes "$TMP/CHANGELOG.md" "0.2.3")
assert_contains "0.2.3 captures Miscellaneous section" "release 0.2.3" "$out"

# 4. Pre-release version
out=$("$PARSE" changelog-notes "$TMP/CHANGELOG.md" "0.2.0-beta.1")
assert_contains "pre-release 0.2.0-beta.1 captured" "beta feature" "$out"

# 5. A version that only appears in a compare URL must not match a heading
out=$("$PARSE" changelog-notes "$TMP/CHANGELOG.md" "0.2.39")
assert_eq "URL-only version falls back to default" "Release 0.2.39" "$out"

# 6. Missing version falls back
out=$("$PARSE" changelog-notes "$TMP/CHANGELOG.md" "9.9.9")
assert_eq "missing version falls back" "Release 9.9.9" "$out"

# 7. GitHub markdown alerts must be preserved verbatim
cat >"$TMP/ALERTS.md" <<'EOF'
# Changelog

## [1.0.0](https://github.com/foo/bar/compare/v0.9.0...v1.0.0) (2026-07-01)

> [!WARNING]
> Breaking change: API endpoint renamed.

> [!IMPORTANT]
> You must run migrations before upgrading.

> [!TIP]
> See the migration guide for details.

> [!CAUTION]
> Do not skip the manual step in section 4.

> [!NOTE]
> This release was tested on Linux, macOS, and Windows.

### Added

* new shiny feature

## [0.9.0](https://github.com/foo/bar/compare/v0.8.0...v0.9.0) (2026-06-15)

prior content
EOF

out=$("$PARSE" changelog-notes "$TMP/ALERTS.md" "1.0.0")
for kind in WARNING IMPORTANT TIP CAUTION NOTE; do
  assert_contains "alert [!${kind}] preserved" "[!${kind}]" "$out"
done
assert_contains "alert body line preserved" "Breaking change: API endpoint renamed." "$out"
assert_contains "trailing section after alerts preserved" "new shiny feature" "$out"
if [[ "$out" == *"prior content"* ]]; then
  echo "  FAIL - capture leaked into 0.9.0 section"
  FAIL=$((FAIL + 1))
else
  echo "  PASS - capture stops cleanly at next version heading"
  PASS=$((PASS + 1))
fi

#####################################################################
# explicit / pre-release version handling
#   Explicit bumps (e.g. 1.5.8-beta.1) are released verbatim, so manifest-diff
#   and changelog-notes must handle pre-release identifiers without colliding
#   with the matching stable version.
#####################################################################
echo
echo "=== explicit / pre-release versions ==="

# 11. Pre-release identifier bump (beta.1 -> beta.2)
out=$(diff_run \
  '{"pkg/go":"1.5.8-beta.2"}' \
  '{"pkg/go":"1.5.8-beta.1"}')
assert_eq "pre-release identifier bump (beta.1 -> beta.2)" \
  '[{"component":"pkg/go","version":"1.5.8-beta.2","tag_name":"pkg/go/v1.5.8-beta.2"}]' \
  "$out"

# 12. Stable to pre-release of the next version
out=$(diff_run \
  '{"pkg/js":"1.5.8-beta.1"}' \
  '{"pkg/js":"1.5.7"}')
assert_eq "stable -> pre-release of next version" \
  '[{"component":"pkg/js","version":"1.5.8-beta.1","tag_name":"pkg/js/v1.5.8-beta.1"}]' \
  "$out"

# 13. Numeric ordering is irrelevant (string compare): beta.9 -> beta.10
out=$(diff_run \
  '{"pkg/java":"2.0.0-beta.10"}' \
  '{"pkg/java":"2.0.0-beta.9"}')
assert_eq "beta.9 -> beta.10 detected (string compare)" \
  '[{"component":"pkg/java","version":"2.0.0-beta.10","tag_name":"pkg/java/v2.0.0-beta.10"}]' \
  "$out"

# 14. Same pre-release version unchanged; no release
echo '{"pkg/go":"3.0.0-rc.1"}' >"$TMP/cur.json"
cp "$TMP/cur.json" "$TMP/prev.json"
"$PARSE" manifest-diff "$TMP/cur.json" "$TMP/prev.json" >/dev/null 2>&1
code=$?
assert_exit_code "unchanged pre-release does not trigger release" 1 "$code"

cat >"$TMP/PRE.md" <<'EOF'
# Changelog

## [1.5.8](https://github.com/foo/bar/compare/v1.5.8-beta.2...v1.5.8) (2026-06-04)

### Fixed

* stable fix for 1.5.8

## [1.5.8-beta.2](https://github.com/foo/bar/compare/v1.5.8-beta.1...v1.5.8-beta.2) (2026-06-03)

### Added

* second beta feature

## [1.5.8-beta.1](https://github.com/foo/bar/compare/v1.5.7...v1.5.8-beta.1) (2026-06-02)

### Added

* first beta feature

## [1.0.0-alpha.10](https://github.com/foo/bar/compare/v1.0.0-alpha.9...v1.0.0-alpha.10) (2026-05-01)

### Added

* alpha ten feature
EOF

# 15. Exact pre-release match captures its own body, no leak from newer beta
out=$("$PARSE" changelog-notes "$TMP/PRE.md" "1.5.8-beta.1")
assert_contains "beta.1 captures its own body" "first beta feature" "$out"
if [[ "$out" == *"second beta feature"* ]]; then
  echo "  FAIL - beta.1 leaked beta.2 content"
  FAIL=$((FAIL + 1))
else
  echo "  PASS - beta.1 does not leak beta.2 content"
  PASS=$((PASS + 1))
fi

# 16. Newer beta captures its own body, stops before older beta
out=$("$PARSE" changelog-notes "$TMP/PRE.md" "1.5.8-beta.2")
assert_contains "beta.2 captures its own body" "second beta feature" "$out"
if [[ "$out" == *"first beta feature"* ]]; then
  echo "  FAIL - beta.2 leaked beta.1 content"
  FAIL=$((FAIL + 1))
else
  echo "  PASS - beta.2 does not leak beta.1 content"
  PASS=$((PASS + 1))
fi

# 17. Stable 1.5.8 must not match the 1.5.8-beta.* headings (boundary on '-')
out=$("$PARSE" changelog-notes "$TMP/PRE.md" "1.5.8")
assert_contains "stable 1.5.8 captures stable body" "stable fix for 1.5.8" "$out"
if [[ "$out" == *"beta feature"* ]]; then
  echo "  FAIL - stable 1.5.8 incorrectly captured a beta section"
  FAIL=$((FAIL + 1))
else
  echo "  PASS - stable 1.5.8 does not capture beta sections"
  PASS=$((PASS + 1))
fi

# 18. Double-digit pre-release identifier (alpha.10)
out=$("$PARSE" changelog-notes "$TMP/PRE.md" "1.0.0-alpha.10")
assert_contains "alpha.10 captured exactly" "alpha ten feature" "$out"

# 19. Partial pre-release prefix must not match (1.5.8-beta -> fallback)
out=$("$PARSE" changelog-notes "$TMP/PRE.md" "1.5.8-beta")
assert_eq "partial pre-release prefix falls back" "Release 1.5.8-beta" "$out"

#####################################################################
# summary
#####################################################################
echo
echo "================================"
echo "  Passed: $PASS"
echo "  Failed: $FAIL"
echo "================================"

if [[ "$FAIL" -gt 0 ]]; then
  exit 1
fi