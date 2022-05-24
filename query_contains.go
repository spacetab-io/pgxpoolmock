package pgxpoolmock

import (
	"fmt"
	"regexp"

	"github.com/golang/mock/gomock"
)

// QueryContainsMatcher implements the gomock matcher interface
// (https://pkg.go.dev/github.com/golang/mock/gomock#Matcher) to match a string (SQL query) that
// contains the given string. This is useful for passing into pgxpoolmock.MockPgxIface as the
// SQL query string argument since SQLC generated code will contain the associated function
// name as well.
//
// Usage: pgxMock.EXPECT().QueryRow(gomock.Any(), QueryContains("GetSomething")).Return(NewRow(1, "foo"))
type QueryContainsMatcher struct{ re *regexp.Regexp }

func QueryContains(s string) gomock.Matcher {
	return &QueryContainsMatcher{
		re: regexp.MustCompile(fmt.Sprintf(".*%s.*", s)),
	}
}

func (m *QueryContainsMatcher) Matches(x interface{}) bool {
	return m.re.MatchString(x.(string))
}

func (m *QueryContainsMatcher) String() string {
	return fmt.Sprintf("matches the regexp `%s`", m.re.String())
}
