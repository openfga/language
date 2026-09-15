#!/usr/bin/env bash
# Tests for downstream-bump.sh
# Run: bash .github/workflows/scripts/downstream-bump.test.sh
#
# The java path is pure text substitution and is exercised end-to-end. The js
# and go paths need a toolchain + network to actually resolve a version, so we
# only assert their hermetic behaviour here: version normalisation and the
# guard/error paths (missing files, missing dependency).

set -uo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
BUMP="$SCRIPT_DIR/downstream-bump.sh"

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

assert_not_contains() {
  local name="$1" needle="$2" haystack="$3"
  if [[ "$haystack" != *"$needle"* ]]; then
    echo "  PASS - $name"
    PASS=$((PASS + 1))
  else
    echo "  FAIL - $name"
    echo "     expected NOT to contain: $needle"
    echo "     actual:                  $haystack"
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

# Run the script inside a throwaway dir, echoing its output. Usage:
#   out=$(run_in <dir> <args...>); code=$?
run_in() {
  local dir="$1"; shift
  ( cd "$dir" && bash "$BUMP" "$@" )
}

fresh() { mktemp -d; }

#####################################################################
# java: full end-to-end text substitution
#####################################################################
echo
echo "=== java ==="

# 1. beta coordinate -> stable, only the language line changes
d=$(fresh)
cat >"$d/build.gradle.kts" <<'EOF'
dependencies {
    implementation("org.apache.commons:commons-lang3:3.20.0")
    implementation("dev.openfga:openfga-language:0.2.0-beta.2")
    implementation("com.diffplug.spotless:spotless-plugin-gradle:8.8.0")
}
EOF
run_in "$d" java 0.2.1 >/dev/null 2>&1
assert_contains "java bumps language coordinate" \
  'dev.openfga:openfga-language:0.2.1' "$(cat "$d/build.gradle.kts")"
assert_contains "java leaves other deps untouched (commons)" \
  'commons-lang3:3.20.0' "$(cat "$d/build.gradle.kts")"
assert_contains "java leaves other deps untouched (spotless)" \
  'spotless-plugin-gradle:8.8.0' "$(cat "$d/build.gradle.kts")"
assert_not_contains "old java version gone" \
  '0.2.0-beta.2' "$(cat "$d/build.gradle.kts")"
rm -rf "$d"

# 2. version normalisation: full tag ref is accepted
d=$(fresh)
echo 'implementation("dev.openfga:openfga-language:0.2.0")' >"$d/build.gradle.kts"
run_in "$d" java pkg/java/v0.3.1 >/dev/null 2>&1
assert_contains "java accepts pkg/java/vX.Y.Z tag ref" \
  'dev.openfga:openfga-language:0.3.1' "$(cat "$d/build.gradle.kts")"
rm -rf "$d"

# 3. leading-v normalisation
d=$(fresh)
echo 'implementation("dev.openfga:openfga-language:0.2.0")' >"$d/build.gradle.kts"
run_in "$d" java v0.4.0 >/dev/null 2>&1
assert_contains "java strips leading v" \
  'dev.openfga:openfga-language:0.4.0' "$(cat "$d/build.gradle.kts")"
rm -rf "$d"

# 4. groovy build.gradle (no .kts) is picked up
d=$(fresh)
echo 'implementation "dev.openfga:openfga-language:0.2.0"' >"$d/build.gradle"
run_in "$d" java 0.2.1 >/dev/null 2>&1
assert_contains "java falls back to build.gradle" \
  'dev.openfga:openfga-language:0.2.1' "$(cat "$d/build.gradle")"
rm -rf "$d"

# 4b. groovy build.gradle with single quotes keeps its closing quote intact
d=$(fresh)
echo "implementation 'dev.openfga:openfga-language:0.2.0'" >"$d/build.gradle"
run_in "$d" java 0.2.1 >/dev/null 2>&1
assert_contains "java handles single-quoted build.gradle" \
  "'dev.openfga:openfga-language:0.2.1'" "$(cat "$d/build.gradle")"
rm -rf "$d"

# 5. missing build file -> error exit
d=$(fresh)
run_in "$d" java 0.2.1 >/dev/null 2>&1
assert_exit_code "java missing build file fails" 1 $?
rm -rf "$d"

# 6. build file without the language dep -> error exit
d=$(fresh)
echo 'implementation("org.other:thing:1.0.0")' >"$d/build.gradle.kts"
run_in "$d" java 0.2.1 >/dev/null 2>&1
assert_exit_code "java without language dep fails" 1 $?
rm -rf "$d"

#####################################################################
# js: guard/error paths (npm resolution needs network, not tested here)
#####################################################################
echo
echo "=== js (guards) ==="

# 7. no workspace declares the dep -> error before touching npm
d=$(fresh)
mkdir -p "$d/client"
echo '{"name":"client","dependencies":{"other":"^1.0.0"}}' >"$d/client/package.json"
err=$(run_in "$d" js 0.2.2 2>&1); code=$?
assert_exit_code "js no matching workspace fails" 1 "$code"
assert_contains "js error mentions syntax-transformer" \
  "syntax-transformer" "$err"
rm -rf "$d"

# 8. no package.json at all -> error
d=$(fresh)
run_in "$d" js 0.2.2 >/dev/null 2>&1
assert_exit_code "js no package.json fails" 1 $?
rm -rf "$d"

#####################################################################
# go: guard/error paths (go get needs network, not tested here)
#####################################################################
echo
echo "=== go (guards) ==="

# 9. no go.mod -> error
d=$(fresh)
run_in "$d" go 0.3.1 >/dev/null 2>&1
assert_exit_code "go no go.mod fails" 1 $?
rm -rf "$d"

# 10. go.mod without the language module -> error before touching go
d=$(fresh)
cat >"$d/go.mod" <<'EOF'
module example.com/foo

go 1.25.0

require github.com/other/thing v1.0.0
EOF
err=$(run_in "$d" go 0.3.1 2>&1); code=$?
assert_exit_code "go without language module fails" 1 "$code"
assert_contains "go error mentions the module path" \
  "github.com/openfga/language/pkg/go" "$err"
rm -rf "$d"

#####################################################################
# generic
#####################################################################
echo
echo "=== generic ==="

# 11. unknown ecosystem -> error
d=$(fresh)
err=$(run_in "$d" rust 1.0.0 2>&1); code=$?
assert_exit_code "unknown ecosystem fails" 1 "$code"
assert_contains "unknown ecosystem message" "Unknown ecosystem" "$err"
rm -rf "$d"

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
