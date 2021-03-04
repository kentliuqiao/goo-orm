package gooorm

import (
	"database/sql"
	"gooorm/dialect"
	"gooorm/log"
	"gooorm/session"
)

type Engine struct {
	db *sql.DB
	dialect dialect.Dialect
}

func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}

	// send a ping to make sure the database connection is alive
	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}
	dial, ok := dialect.GetDialect(driver)
	if !ok {
		log.Error("dialect %s not found", driver)
		return
	}

	e = &Engine{db: db, dialect: dial}
	log.Info("connect database succeed")
	return
}

func (e *Engine) Close() {
	if err := e.db.Close(); err != nil {
		log.Error(err)
	}

	log.Info("close database succeed")
}

func (e *Engine) NewSession() *session.Session {
	return session.New(e.db, e.dialect)
}
