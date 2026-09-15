#!/bin/sh
# Modified by Oh My Sub2API contributors on 2026-09-15; see CHANGES.md.
set -eu

repo_root=$(CDPATH= cd -- "$(dirname -- "$0")/../.." && pwd)
cd "$repo_root"

fail() {
  printf 'docker runtime resources test failed: %s\n' "$1" >&2
  exit 1
}

assert_line() {
  file=$1
  line=$2
  grep -Fqx "$line" "$file" || fail "$file is missing: $line"
}

assert_count() {
  file=$1
  line=$2
  expected=$3
  actual=$(grep -Fxc "$line" "$file" || true)
  [ "$actual" -eq "$expected" ] || fail "$file has $actual occurrences of '$line', expected $expected"
}

test -s backend/resources/model-pricing/model_prices_and_context_window.json || \
  fail 'fallback pricing data is missing or empty'

assert_line Dockerfile.goreleaser 'COPY --chown=sub2api:sub2api backend/resources /app/resources'
assert_line deploy/Dockerfile 'COPY --from=backend-builder --chown=sub2api:sub2api /app/backend/resources /app/resources'
assert_count .goreleaser.yaml '      - backend/resources' 2
assert_count .goreleaser.yaml '      - LICENSE' 3
assert_count .goreleaser.yaml '      - COPYING' 3
assert_count .goreleaser.yaml '      - NOTICE' 3
for dockerfile in Dockerfile Dockerfile.goreleaser deploy/Dockerfile; do
  assert_line "$dockerfile" 'COPY LICENSE COPYING NOTICE /usr/share/licenses/oh-my-sub2api/'
done

printf 'docker runtime resources test passed\n'
