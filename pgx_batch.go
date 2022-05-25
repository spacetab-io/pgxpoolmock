package pgxpoolmock

import (
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

// BatchResults is the same interface as pgx.BatchResults, placed here for mocking.
// https://github.com/jackc/pgx/blob/dc0ad04ff58f72f4819289f54745a36124cdbec3/batch.go#L35-L52
type BatchResults interface {
	Exec() (pgconn.CommandTag, error)
	Query() (pgx.Rows, error)
	QueryRow() pgx.Row
	QueryFunc(scans []interface{}, f func(pgx.QueryFuncRow) error) (pgconn.CommandTag, error)
	Close() error
}
