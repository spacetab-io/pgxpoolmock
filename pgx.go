package pgxpoolmock

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

// Implements pgxpool.PgxPool and pgx.Tx
type PgxIface interface {
	// pgx.Tx
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnName []string, rowSvc pgx.CopyFromSource) (int64, error)
	LargeObjects() pgx.LargeObjects
	Prepare(ctx context.Context, name, sql string) (*pgconn.StatementDescription, error)
	Conn() *pgx.Conn

	// pgx.PgxPool
	Close()
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	QueryFunc(ctx context.Context, sql string, args []interface{}, scans []interface{}, f func(pgx.QueryFuncRow) error) (pgconn.CommandTag, error)
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
	BeginFunc(ctx context.Context, f func(pgx.Tx) error) error
	BeginTxFunc(ctx context.Context, txOptions pgx.TxOptions, f func(pgx.Tx) error) error
}
