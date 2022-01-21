package stores

import (
	"database/sql"
	"log"
	reflect "reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/zopping/mock-test/models"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}
func Test_user(t *testing.T) {
	db, mock := NewMock()
	s := New(db)
	query := "select id, name from users where id =? "
	tcs := []struct {
		testCase    int
		id          int
		expectedErr error
		expectedOut *models.User
	}{
		{
			testCase:    1,
			id:          1,
			expectedErr: nil,
			expectedOut: &models.User{Id: 1, Name: "test"},
		},
	}
	for _, tc := range tcs {
		mock.ExpectQuery(query).WithArgs(tc.id).
			WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "test"))
		resp, err := s.Find(tc.id)
		if !reflect.DeepEqual(resp, tc.expectedOut) {
			t.Errorf("TestCase[%v] Expected: \t%v\nGot: \t%v\n", tc.testCase, tc.expectedOut, resp)
		}
		if !reflect.DeepEqual(err, tc.expectedErr) {
			t.Errorf("TestCase[%v] Expected: \t%v\nGot: \t%v\n", tc.testCase, tc.expectedErr, err)
		}
	}
}

func Test_Create(t *testing.T) {
	db, mock := NewMock()
	s := New(db)
	query := "insert into users(id, name) values (?, ?)"

	tests := []struct {
		testCase int
		userName string
		userId   int
		resp     int
		err      error
		mock     []interface{}
	}{
		{
			testCase: 1,
			userName: "test",
			userId:   1,
			resp:     1,
			err:      nil,
			mock: []interface{}{
				mock.ExpectExec(query).WithArgs(1, "test").WillReturnResult(sqlmock.NewResult(1, 1)),
			},
		},
	}
	for _, tc := range tests {
		resp, err := s.Create(tc.userId, tc.userName)
		if !reflect.DeepEqual(resp, tc.resp) {
			t.Errorf("TestCase[%v] Expected: \t%v\nGot: \t%v\n", tc.testCase, tc.resp, resp)
		}
		if !reflect.DeepEqual(err, tc.err) {
			t.Errorf("TestCase[%v] Expected: \t%v\nGot: \t%v\n", tc.testCase, tc.err, err)
		}
	}

}

func Test_Update(t *testing.T) {
	db, mock := NewMock()
	s := New(db)
	query := "update users set name = ? where id = ? "

	tests := []struct {
		testCase int
		userName string
		userId   int
		resp     int
		err      error
		mock     []interface{}
	}{
		{
			testCase: 1,
			userName: "test_new",
			userId:   1,
			resp:     1,
			err:      nil,
			mock: []interface{}{
				mock.ExpectExec(query).WithArgs("test_new", 1).WillReturnResult(sqlmock.NewResult(1, 1)),
			},
		},
	}
	for _, tc := range tests {
		err := s.Update(tc.userId, tc.userName)
		if !reflect.DeepEqual(err, tc.err) {
			t.Errorf("TestCase[%v] Expected: \t%v\nGot: \t%v\n", tc.testCase, tc.err, err)
		}
	}

}
