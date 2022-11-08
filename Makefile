.PHONY: mock
mock: bin/mockgen
	@mockgen \
		-source=pgx.go \
		-destination=pgx_mock.go \
		-package=pgxpoolmock
	@mockgen \
		-source=pgx_batch.go \
		-destination=pgx_batch_mock.go \
		-package=pgxpoolmock

.PHONY: sql
sql: bin/sqlc
	@sqlc generate

MOCKGEN_VERSION?=1.6.0
bin/mockgen:
	GOBIN=$(shell pwd)/bin/ \
		go install github.com/golang/mock/mockgen@v${MOCKGEN_VERSION}

# sqlc generates type-safe code from SQL.
# https://github.com/kyleconroy/sqlc
SQLC_VERSION?=1.15.0
bin/sqlc:
	GOBIN=$(shell pwd)/bin/ \
		go install github.com/kyleconroy/sqlc/cmd/sqlc@v${SQLC_VERSION}


# ----
## LINTER stuff start

linter_include_check:
	@[ -f linter.mk ] && echo "linter.mk include exists" || (echo "getting linter.mk from github.com" && curl -sO https://raw.githubusercontent.com/spacetab-io/makefiles/master/golang/linter.mk)

.PHONY: lint
lint: linter_include_check
	@make -f linter.mk go_lint

## LINTER stuff end
# ----


# ----
## TESTS stuff start

tests_include_check:
	@[ -f tests.mk ] && echo "tests.mk include exists" || (echo "getting tests.mk from github.com" && curl -sO https://raw.githubusercontent.com/spacetab-io/makefiles/master/golang/tests.mk)

tests: tests_include_check
	@make -f tests.mk go_tests
.PHONY: tests

tests_html: tests_include_check
	@make -f tests.mk go_tests_html
.PHONY: tests_html

## TESTS stuff end
# ----
