package main

import (
	"gooorm"
	"gooorm/log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	engine, err := gooorm.NewEngine("sqlite3", "goo.db")
	if err != nil {
		log.Error(err)
	}
	defer engine.Close()

	s := engine.NewSession()
	_, _ = s.Raw("drop table if exists user;").Exec()
	_, _ = s.Raw("create table user(name text);").Exec()
	_, _ = s.Raw("create table user(name text);").Exec()
	res, _ := s.Raw("insert into user(`name`) values(?),(?)", "Tom", "Jerry").Exec()
	count, _ := res.RowsAffected()
	log.Infof("Exec success, %d affected\n", count)
}
