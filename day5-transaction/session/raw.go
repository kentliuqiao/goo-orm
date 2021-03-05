package session

import (
	"database/sql"
	"gooorm/clause"
	"gooorm/dialect"
	"gooorm/log"
	"gooorm/schema"
	"strings"
)

type Session struct {
	db      *sql.DB // db connection
	tx *sql.Tx
	dialect dialect.Dialect
	refTable *schema.Schema // to store parsed struct
	clause clause.Clause
	sql     strings.Builder // unprepared sql
	sqlVars []interface{}
}

type CommonDB interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
}

var _ CommonDB = (*sql.DB)(nil)
var _ CommonDB = (*sql.Tx)(nil)

func New(db *sql.DB, dialect dialect.Dialect) *Session {
	return &Session{
		db: db,
		dialect: dialect,
	}
}

func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVars = nil
	s.clause = clause.Clause{}
}

func (s *Session) DB() CommonDB {
	if s.tx != nil {
		return s.tx
	}

	return s.db
}

func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlVars = append(s.sqlVars, values...)

	return s
}

func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear()

	log.Info(s.sql.String(), s.sqlVars)

	if result, err = s.db.Exec(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}

	return
}

func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()

	log.Info(s.sql.String(), s.sqlVars)

	return s.db.QueryRow(s.sql.String(), s.sqlVars...)
}

func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()

	log.Info(s.sql.String(), s.sqlVars)

	if rows, err = s.db.Query(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}

	return
}
