package session

import "testing"

type User struct {
	Name string `gooorm:"PRIMARY KEY"`
	Age int
}

func TestSession_CreateTable(t *testing.T) {
	s := NewSession().Model(&User{})
	s.DropTable()
	s.CreateTable()
	if !s.HasTable() {
		t.Error("failed to create table User")
	}
}

func TestSession_Model(t *testing.T) {
	s := NewSession().Model(&User{})
	table := s.RefTable()
	s.Model(&Session{})
	if s.RefTable().Name != "Session" || table.Name != "User" {
		t.Error("failed to change model")
	}
}
