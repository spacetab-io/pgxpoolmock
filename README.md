### PGX POOL MOCK

This repo provides a mock of `pgxpool.Pool`, `pgx.Tx`, and `pgx.BatchResult` from https://github.com/jackc/pgx/v5 so that you can test your data access code locally without using a real database.

It was forked from https://www.github.com/chrisyxlee/pgxpoolmock to add support for transactions and batch sends with `pgxpool.Pool` and `pgx/v5`.

### How to install

```
go get -u github.com/spacetab-io/pgxpoolmock
```

### How to Use

See file `order_dao_example_test.go` to figure out how to use it. Or see the below:

```go
package pgxpoolmock_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/spacetab-io/pgxpoolmock"
	"github.com/spacetab-io/pgxpoolmock/sqlc"
	"github.com/spacetab-io/pgxpoolmock/testdata"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	mockPool := pgxpoolmock.NewMockPgxIface(ctrl)
	columns := []string{"id", "price"}
	pgxRows := pgxpoolmock.NewRows(columns).AddRow(100, 100000.9).ToPgxRows()
	mockPool.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRows, nil)
	orderDao := testdata.OrderDAO{
		Pool: mockPool,
	}

	// when
	actualOrder := orderDao.GetOrderByID(1)

	// then
	assert.NotNil(t, actualOrder)
	assert.Equal(t, 100, actualOrder.ID)
	assert.Equal(t, 100000.9, actualOrder.Price)
}

func TestMap(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	mockPool := pgxpoolmock.NewMockPgxIface(ctrl)
	columns := []string{"id", "price"}
	pgxRows := pgxpoolmock.NewRows(columns).AddRow(100, 100000.9).ToPgxRows()
	assert.NotEqual(t, "with empty rows", fmt.Sprintf("%s", pgxRows))

	mockPool.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRows, nil)
	orderDao := testdata.OrderDAO{
		Pool: mockPool,
	}

	// when
	actualOrder := orderDao.GetOrderMapByID(1)

	// then
	assert.NotNil(t, actualOrder)
	assert.Equal(t, 100, actualOrder["ID"])
	assert.Equal(t, 100000.9, actualOrder["Price"])
}

func TestTx(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockPool := pgxpoolmock.NewMockPgxIface(ctrl)

	// begin tx - given
	mockPool.EXPECT().BeginTx(gomock.Any(), gomock.Any()).Return(mockPool, nil)
	tx, err := mockPool.BeginTx(context.Background(), pgx.TxOptions{})
	assert.NoError(t, err)

	mockPool.EXPECT().QueryRow(gomock.Any(), "blah").Return(
		pgxpoolmock.NewRow("1"))
	row := tx.QueryRow(context.Background(), "blah")
	var s string
	err = row.Scan(&s)
	assert.NoError(t, err)
	assert.EqualValues(t, s, "1")

	mockPool.EXPECT().Rollback(gomock.Any()).Return(nil)
	assert.NoError(t, tx.Rollback(context.Background()))
}

func TestQueryMatcher(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockPool := pgxpoolmock.NewMockPgxIface(ctrl)

	mockPool.EXPECT().QueryRow(gomock.Any(), pgxpoolmock.QueryContains("blah")).Return(
		pgxpoolmock.NewRow("1"))
	row := mockPool.QueryRow(context.Background(), "SELECT blah FROM some_table;")
	var s string
	err := row.Scan(&s)
	assert.NoError(t, err)
}

func TestBatchResults(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockPool := pgxpoolmock.NewMockPgxIface(ctrl)
	mockBatch := pgxpoolmock.NewMockBatchResults(ctrl)

	// Unfortunately pgx.Batch isn't passed as an interface, so there's no way to mock the
	// intermediate data. The best bet is probably to separate the data creation and handling
	// into a separate function to test separately.
	mockPool.EXPECT().SendBatch(gomock.Any(), gomock.Any()).Return(mockBatch)
	mockBatch.EXPECT().QueryRow().Return(pgxpoolmock.NewRow(int32(1)))
	mockBatch.EXPECT().QueryRow().Return(pgxpoolmock.NewRow(int32(2)))
	mockBatch.EXPECT().QueryRow().Return(pgxpoolmock.NewRow(int32(3)))
	mockBatch.EXPECT().QueryRow().Return(pgxpoolmock.NewRow(int32(0)).WithError(pgxpoolmock.ErrEndBatchResult))

	q := sqlc.New(mockPool)
	var inserted int32
	q.InsertAuthors(context.Background(), []string{"a", "b", "c"}).QueryRow(func(i int, authorID int32, err error) {
		inserted++
		assert.Equal(t, inserted, authorID)
		assert.NoError(t, err)
	})
	assert.Equal(t, int32(3), inserted)
}
```
### How to test

    make tests


### Check code style

    make lint