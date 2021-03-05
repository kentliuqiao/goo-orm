package session

import (
	"database/sql"
	"gooorm/dialect"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

var (
	TestDB *sql.DB
	TestDialect, _ = dialect.GetDialect("sqlite3")
)

func TestMain(m *testing.M) {
	TestDB, _ = sql.Open("sqlite3", "../goo.db")

	code := m.Run()

	_ = TestDB.Close()
	os.Exit(code)
}

func NewSession() *Session {
	return New(TestDB, TestDialect)
}

func TestSession_Exec(t *testing.T) {
	s := NewSession()
	s.Raw("DROP TABLE IF EXISTS User;").Exec()
	s.Raw("CREATE TABLE User(Name text);").Exec()
	res, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Jerry").Exec()
	if count, err := res.RowsAffected(); err != nil || count != 2 {
		t.Error("expect 2 but got ", count)
	}
}

func TestSession_QueryRows(t *testing.T) {
	s := NewSession()
	s.Raw("DROP TABLE IF EXISTS User;").Exec()
	s.Raw("CREATE TABLE User(Name text);").Exec()
	row := s.Raw("SELECT count(*) FROM User").QueryRow()
	var count int
	if err := row.Scan(&count); err != nil || count != 0 {
		t.Error("failed to query db", err)
	}
}
