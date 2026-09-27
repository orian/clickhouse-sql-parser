
PROGRAM=clickhouse-sql-parser
PKG_FILES=`go list ./... | sed -e 's=github.com/orian/clickhouse-sql-parser/=./='`

CCCOLOR="\033[37;1m"
MAKECOLOR="\033[32;1m"
ENDCOLOR="\033[0m"

all: $(PROGRAM)

.PHONY: all

$(PROGRAM):
	go build -o $(PROGRAM) main.go

test:
	@go test -v ./... -covermode=atomic -coverprofile=coverage.out -race -compatible

update_test:
	@go test -v ./... -update -race -compatible

lint:
	@printf $(CCCOLOR)"GolangCI Lint...\n"$(ENDCOLOR)
	@golangci-lint run --timeout 20m0s

.PHONY: hooks fmt-check

# hooks enables the repository's git hooks (.githooks), e.g. the gofmt
# pre-commit check.
hooks:
	git config core.hooksPath .githooks

# fmt-check fails if any Go file is not gofmt-formatted.
fmt-check:
	@test -z "$$(gofmt -l $$(git ls-files '*.go'))" || { gofmt -l $$(git ls-files '*.go'); exit 1; }

.PHONY: docs-sql-extract docs-sql-test docs-sql-update docs-sql-strict

docs-sql-extract:
	python3 scripts/extract_docs_sql.py

docs-sql-test:
	go test ./parser -run '^TestDocumentationSQL$$' -docs-sql -count=1 -timeout 3m

docs-sql-update:
	go test ./parser -run '^TestDocumentationSQL$$' -docs-sql-update -count=1 -timeout 3m

docs-sql-strict:
	go test ./parser -run '^TestDocumentationSQL$$' -docs-sql-strict -count=1 -timeout 3m

.PHONY: clickhouse-sql-extract clickhouse-sql-test clickhouse-sql-update clickhouse-sql-strict

clickhouse-sql-extract:
	python3 scripts/extract_clickhouse_sql.py

clickhouse-sql-test:
	go test ./parser -run '^TestClickHouseSQL$$' -clickhouse-sql -count=1 -timeout 3m

clickhouse-sql-update:
	go test ./parser -run '^TestClickHouseSQL$$' -clickhouse-sql-update -count=1 -timeout 3m

clickhouse-sql-strict:
	go test ./parser -run '^TestClickHouseSQL$$' -clickhouse-sql-strict -count=1 -timeout 3m

FUZZ_TIME ?= 30s
FUZZ_PARALLEL ?= 4
FUZZ_TIMEOUT ?= 5m
FUZZ_TARGET ?= FuzzParseStmts
FUZZ_SQL_CORPORA ?=
.PHONY: fuzz
fuzz:
	go test ./parser -run '^$$' -fuzz '^$(FUZZ_TARGET)$$' -fuzztime $(FUZZ_TIME) -parallel $(FUZZ_PARALLEL) -timeout $(FUZZ_TIMEOUT) -fuzz-sql-corpora '$(FUZZ_SQL_CORPORA)'
