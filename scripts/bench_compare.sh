#!/usr/bin/env bash
# Compare the parser benchmarks of two git refs on identical inputs.
#
#   scripts/bench_compare.sh BASE [HEAD]      e.g. scripts/bench_compare.sh v1.0.4
#
# BASE and HEAD are checked out into temporary worktrees (HEAD defaults to the
# current HEAD). Both are measured with the current parser/benchmark_test.go,
# on the current testdata/query and testdata/benchdata fixtures that parse in
# both refs (the check uses each ref's CLI). Rounds alternate between the refs
# so machine noise affects both alike, and benchstat compares the results.
#
# Environment: ROUNDS (default 10), BENCH (benchmark regexp, default
# 'Corpus|Huge'), BENCHTIME (default 1s).
set -euo pipefail

if [ $# -lt 1 ] || [ $# -gt 2 ]; then
	echo "usage: $0 BASE [HEAD]" >&2
	exit 2
fi
base=$1
head=${2:-HEAD}
rounds=${ROUNDS:-10}
bench=${BENCH:-Corpus|Huge}
benchtime=${BENCHTIME:-1s}

repo=$(git rev-parse --show-toplevel)
work=$(mktemp -d)
cleanup() {
	git -C "$repo" worktree remove --force "$work/base" 2>/dev/null || true
	git -C "$repo" worktree remove --force "$work/head" 2>/dev/null || true
	rm -rf "$work"
}
trap cleanup EXIT

git -C "$repo" worktree add -q --detach "$work/base" "$base"
git -C "$repo" worktree add -q --detach "$work/head" "$head"

# Freeze the inputs: the current fixtures.
mkdir "$work/inputs"
cp "$repo"/parser/testdata/query/*.sql "$repo"/parser/testdata/benchdata/*.sql "$work/inputs/"

for side in base head; do
	cp "$repo/parser/benchmark_test.go" "$work/$side/parser/benchmark_test.go"
	(cd "$work/$side" && go build -o "$work/$side.cli" . && go test -c -o "$work/$side.test" ./parser)
done

# Keep the inputs both refs parse.
: >"$work/inputs.txt"
skipped=0
for file in "$work"/inputs/*.sql; do
	if "$work/base.cli" -format -f "$file" >/dev/null 2>&1 &&
		"$work/head.cli" -format -f "$file" >/dev/null 2>&1; then
		echo "$file" >>"$work/inputs.txt"
	else
		skipped=$((skipped + 1))
	fi
done
echo "inputs: $(wc -l <"$work/inputs.txt") files ($skipped skipped: not parsed by both refs)" >&2

for i in $(seq 1 "$rounds"); do
	echo "round $i/$rounds" >&2
	for side in base head; do
		(cd "$work/$side/parser" &&
			BENCH_INPUTS="$work/inputs.txt" "$work/$side.test" \
				-test.run '^$' -test.bench "$bench" -test.benchmem -test.benchtime "$benchtime") >>"$work/$side.txt"
	done
done

echo "base = $base ($(git -C "$repo" rev-parse --short "$base")), head = $head ($(git -C "$repo" rev-parse --short "$head"))" >&2
go run golang.org/x/perf/cmd/benchstat@latest "base=$work/base.txt" "head=$work/head.txt"
