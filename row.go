package pgxpoolmock

import (
	"fmt"
	"reflect"
)

// Row implements the pgx.Row interface and can be passed into pgxpoolmock.PgxPoolMock.QueryRow
// as the expected returned row.
//
// Usage: pgxMock.EXPECT().QueryRow(gomock.Any(), HasString("GetEntityByID")).Return(NewRow(1, "foo"))
type Row struct {
	values []interface{}
}

func NewRow(values ...interface{}) *Row {
	return &Row{values: values}
}

func (r *Row) Scan(dest ...interface{}) error {
	if len(r.values) != len(dest) {
		panic(fmt.Errorf("expected scan to be called with same number of arguments\ngot %d\n%+v\nwant %d\n%+v", len(dest), dest, len(r.values), r.values))
	}

	for i := range dest {
		value := r.values[i]
		valueType := reflect.TypeOf(value).String()
		valueRV := reflect.ValueOf(value)

		if dest[i] == nil {
			panic(fmt.Errorf("unexpected nil value for arg %d, want type %s", i, valueType))
		}

		dstRV := reflect.ValueOf(dest[i])
		if dstRV.Kind() != reflect.Ptr {
			panic(fmt.Errorf("expected scan to be called with pointers: got %s, want %s", reflect.TypeOf(dest[i]).String(), valueType))
		}

		innerDstRV := reflect.Indirect(dstRV)
		dstType := innerDstRV.Type().String()
		if dstType != valueType {
			panic(fmt.Errorf("scan with unexpected arg %d: got type %s, want type %s", i, dstType, valueType))
		}

		innerDstRV.Set(valueRV)
	}
	return nil
}
