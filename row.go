package pgxpoolmock

import (
	"fmt"
	"reflect"
)

// Row implements the pgx.Row interface and can be passed into pgxpoolmock.PgxPoolMock.QueryRow
// as the expected returned row.
//
// Usage: pgxMock.EXPECT().QueryRow(gomock.Any(), HasString("GetEntityByID")).Return(NewRow(1, "foo")).
type Row struct {
	values []interface{}
	err    error
}

func NewRow(values ...any) *Row {
	return &Row{
		values: values,
		err:    nil,
	}
}

func (r *Row) WithError(err error) *Row {
	r.err = err

	return r
}

func (r *Row) Scan(dest ...any) error {
	if len(r.values) != len(dest) {
		return fmt.Errorf(
			"%w\ngot %d\n%+v\nwant %d\n%+v",
			ErrScanExpectedToHaveSameNumberOfArgs,
			len(dest),
			dest,
			len(r.values),
			r.values,
		)
	}

	for i := range dest {
		value := r.values[i]
		valueType := reflect.TypeOf(value).String()
		valueRV := reflect.ValueOf(value)

		if dest[i] == nil {
			return fmt.Errorf("%w %d, want type %s", ErrUnexpectedNilVal, i, valueType)
		}

		dstRV := reflect.ValueOf(dest[i])
		if dstRV.Kind() != reflect.Ptr {
			return fmt.Errorf("%w: got %s, want %s", ErrToBeCalledWithPointers, reflect.TypeOf(dest[i]).String(), valueType)
		}

		innerDstRV := reflect.Indirect(dstRV)

		dstType := innerDstRV.Type().String()
		if dstType != valueType {
			return fmt.Errorf("%w %d: got type %s, want type %s", ErrUnexpectedArg, i, dstType, valueType)
		}

		innerDstRV.Set(valueRV)
	}

	return r.err
}
