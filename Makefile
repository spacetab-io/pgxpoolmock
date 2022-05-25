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

.PHONY: test
test:
	go test -v ./...

MOCKGEN_VERSION?=1.6.0
bin/mockgen:
	GOBIN=$(shell pwd)/bin/ \
		go install github.com/golang/mock/mockgen@v${MOCKGEN_VERSION}

# sqlc generates type-safe code from SQL.
# https://github.com/kyleconroy/sqlc
SQLC_VERSION?=1.13.0
bin/sqlc:
	GOBIN=$(shell pwd)/bin/ \
		go install github.com/kyleconroy/sqlc/cmd/sqlc@v${SQLC_VERSION}
