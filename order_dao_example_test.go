package pgxpoolmock_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/driftprogramming/pgxpoolmock"
	"github.com/driftprogramming/pgxpoolmock/testdata"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v4"
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
	assert.NoError(t, err)
	var s string
	row.Scan(&s)
	assert.EqualValues(t, s, "1")

	mockPool.EXPECT().Rollback(gomock.Any()).Return(nil)
	assert.NoError(t, tx.Rollback(context.Background()))
}
