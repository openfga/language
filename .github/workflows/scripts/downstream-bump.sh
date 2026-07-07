#!/usr/bin/env bash
# Applies a version bump of an openfga-language artifact inside a checked-out
# downstream repository. Kept out of the workflow YAML so the edit logic can be
# unit-tested independently (mirrors parse-release.sh).
#
# Run from the root of the *downstream* repo checkout.
#
# Usage:
#   downstream-bump.sh js   <version>   # bumps @openfga/syntax-transformer in
#                                        # client/ and server/ workspaces
#   downstream-bump.sh java <version>   # bumps dev.openfga:openfga-language in
#                                        # build.gradle.kts
#   downstream-bump.sh go   <version>   # bumps github.com/openfga/language/pkg/go
#                                        # in the module containing go.mod
#
# <version> is the plain semver (e.g. 0.2.2), without a leading "v".
#
# Requires the relevant toolchain on PATH: npm for js, go for go. java needs
# nothing beyond coreutils.
set -euo pipefail

ecosystem="${1:?ecosystem required: js|java|go}"
version="${2:?version required}"

# Normalise: callers may pass a tag ref like "pkg/js/v0.2.2" or "v0.2.2".
version="${version##*/}"   # strip any "pkg/<lang>/" prefix
version="${version#v}"     # strip a leading "v"

case "${ecosystem}" in
  js)
    # Update each workspace that declares the dependency. npm rewrites both
    # package.json (to ^<version>) and package-lock.json (with the correct
    # integrity hash + refreshed transitive deps), matching the manual PRs.
    found=0
    for dir in client server; do
      if [[ -f "${dir}/package.json" ]] \
         && grep -q '@openfga/syntax-transformer' "${dir}/package.json"; then
        echo "Bumping @openfga/syntax-transformer in ${dir}/"
        ( cd "${dir}" && npm install "@openfga/syntax-transformer@^${version}" )
        found=1
      fi
    done
    if [[ "${found}" -eq 0 ]]; then
      echo "No workspace declares @openfga/syntax-transformer" >&2
      exit 1
    fi
    ;;

  java)
    file="build.gradle.kts"
    [[ -f "${file}" ]] || file="build.gradle"
    if [[ ! -f "${file}" ]]; then
      echo "No build.gradle(.kts) found" >&2
      exit 1
    fi
    if ! grep -q 'dev\.openfga:openfga-language:' "${file}"; then
      echo "${file} does not depend on dev.openfga:openfga-language" >&2
      exit 1
    fi
    echo "Bumping dev.openfga:openfga-language to ${version} in ${file}"
    # Replace the version segment of the coordinate, whatever it currently is
    # (e.g. 0.2.0-beta.2 -> 0.2.1). Only touch chars up to the closing quote.
    # perl for portability (GNU vs BSD sed differ on -i / backrefs).
    perl -i -pe "s/(dev\.openfga:openfga-language:)[^\"]+/\${1}${version}/g" "${file}"
    ;;

  go)
    module="github.com/openfga/language/pkg/go"
    if [[ ! -f "go.mod" ]]; then
      echo "No go.mod in current directory" >&2
      exit 1
    fi
    if ! grep -q "${module}" go.mod; then
      echo "go.mod does not require ${module}" >&2
      exit 1
    fi
    echo "Bumping ${module} to v${version}"
    # proxy.golang.org first; fall back to direct if the tag isn't indexed yet.
    export GOPROXY="proxy.golang.org,direct"
    go get "${module}@v${version}"
    go mod tidy
    ;;

  *)
    echo "Unknown ecosystem: ${ecosystem} (expected js|java|go)" >&2
    exit 1
    ;;
esac

echo "Bump applied for ${ecosystem} -> ${version}"
