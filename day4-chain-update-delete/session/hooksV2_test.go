package session

import (
	"gooorm/log"
	"testing"
)

type AccountV2 struct {
	ID int `gooorm:"PRIMARY KEY"`
	Password string
}

func (a *AccountV2) BeforeInsert(s *Session) error {
	log.Info("before insert", a)
	a.ID += 1000
	return nil
}

func (a *AccountV2) AfterQuery(s *Session) error {
	log.Info("after query", a)
	a.Password = "******"
	return nil
}

func TestSession_CallMethodV2(t *testing.T) {
	s := NewSession().Model(&AccountV2{})
	s.DropTable()
	s.CreateTable()
	s.Insert(&AccountV2{1, "123456"}, &AccountV2{23, "asdqwe"})

	u := &AccountV2{}

	err := s.First(u)
	if err != nil || u.ID != 1001 || u.Password != "******" {
		t.Errorf("failed to call hooks after query, got %v", u)
	}
}
