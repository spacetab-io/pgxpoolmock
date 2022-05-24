MOCKGEN_VERSION?=1.6.0
bin/mockgen:
	GOBIN=$(shell pwd)/bin/ \
		go install github.com/golang/mock/mockgen@v${MOCKGEN_VERSION}

.PHONY: mock
mock: bin/mockgen
	@mockgen -source=pgx.go -destination=pgx_mock.go -package=pgxpoolmock
