package database

import (
	"context"
	"ephelsa/my-career/test/database"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

type tableTest struct {
	Id    string
	Value string
}

func TestNewRowsByQueryContext(t *testing.T) {
	tests := []struct {
		description string

		query     string
		mockQuery string
		arg       string
		withArg   bool

		expectedRows *sqlmock.Rows
		expectedLen  int
	}{
		{
			description:  "fetch with args",
			query:        `SELECT id, value FROM test WHERE id = $1`,
			mockQuery:    `SELECT id, value FROM test WHERE id = ?`,
			arg:          "3",
			withArg:      true,
			expectedRows: sqlmock.NewRows([]string{"id", "value"}).AddRow("1", "hey"),
			expectedLen:  1,
		},
		{
			description:  "fetch without args",
			query:        "SELECT id, value FROM test",
			mockQuery:    "SELECT id, value FROM test",
			withArg:      false,
			expectedRows: sqlmock.NewRows([]string{"id", "value"}).AddRow("1", "hey").AddRow("2", "there"),
			expectedLen:  2,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			db, mock := database.NewMockDatabase(t)
			expectedQuery := mock.ExpectQuery(test.mockQuery)
			if test.withArg {
				expectedQuery.WithArgs(test.arg)
			}
			expectedQuery.WillReturnRows(test.expectedRows)

			// DB conn
			rows, err := NewRowsByQueryContext(db, context.Background(), test.query, test.arg)
			// Unexpected error
			assert.NoErrorf(t, err, test.description)

			result := make([]tableTest, 0)
			for rows.Next() {
				r := tableTest{}
				err = rows.Scan(&r.Id, &r.Value)
				// Unexpected error
				assert.NoErrorf(t, err, test.description)

				result = append(result, r)
			}

			assert.Equalf(t, test.expectedLen, len(result), test.description)
		})
	}
}
