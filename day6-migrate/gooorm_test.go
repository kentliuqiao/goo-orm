package gooorm

import (
	"errors"
	"reflect"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"gooorm/session"
)

type User struct {
	Name string `gooorm:"PRIMARY KEY"`
	Age int
}

type Pet struct {
	Name string `gooorm:"PRIMARY KEY"`
	Age int
}

func TestNewEngine(t *testing.T) {
	engine := OpenDB(t)
	defer engine.Close()
}

func OpenDB(t *testing.T) *Engine {
	t.Helper()
	engine, err := NewEngine("sqlite3", "goo.db")
	if err != nil {
		t.Fatal("failed to connect", err)
	}

	return engine
}

func TestEngine_Transaction(t *testing.T) {
	t.Run("rollback", func(t *testing.T) {
		transactionRollBack(t)
	})

	t.Run("commit", func(t *testing.T) {
		transactionCommit(t)
	})
}

func transactionRollBack(t *testing.T) {
	engine := OpenDB(t)
	defer engine.Close()

	s := engine.NewSession()
	s.Model(&User{}).DropTable()
	_, err := engine.Transaction(func(session *session.Session) (res interface{}, err error) {
		_ = s.Model(&User{}).CreateTable()
		_, err = s.Insert(&User{"Tom", 12})
		return nil, errors.New("error")
	})
	if err == nil || s.HasTable() {
		t.Error("failed to rollback")
	}
}

func transactionCommit(t *testing.T) {
	engine := OpenDB(t)
	defer engine.Close()

	s := engine.NewSession()
	s.Model(&User{}).DropTable()
	_, err := engine.Transaction(func(s *session.Session) (res interface{}, err error) {
		s.Model(&User{}).CreateTable()
		_, err = s.Insert(&User{"Tom", 21})
		return
	})
	u := &User{}
	s.First(u)
	if err != nil || u.Name != "Tom" {
		t.Error("failed to commit")
	}
}

func TestEngine_Migrate(t *testing.T) {
	engine := OpenDB(t)
	defer engine.Close()

	s := engine.NewSession()
	s.Raw("DROP TABLE IF EXISTS User;").Exec()
	s.Raw("CREATE TABLE User(Name text PRIMARY KEY, XXX integer);").Exec()
	s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()

	engine.Migrate(&User{})

	rows, _ := s.Raw("SELECT * FROM User").QueryRows()
	defer rows.Close()
	cols, _ := rows.Columns()
	if !reflect.DeepEqual(cols, []string{"Name", "Age"}) {
		t.Error("failed to migrate table User, got columns", cols)
	}
}
