package session

import "testing"

var (
	u1 = &User{"Tom", 12}
	u2 = &User{"Jerry", 33}
	u3 = &User{"Blake", 23}
)

func TestSession_Insert(t *testing.T) {
	s := testRecordInit(t)
	affected, err := s.Insert(u3)
	if err != nil || affected != 1 {
		t.Error("failed to test records")
	}
}

func TestSession_Find(t *testing.T) {
	s := testRecordInit(t)
	var users []User
	if err := s.Find(&users); err != nil || len(users) != 2 {
		t.Error("failed to query all")
	}
}

func testRecordInit(t *testing.T) *Session {
	t.Helper()
	s := NewSession().Model(&User{})
	err1 := s.DropTable()
	err2 := s.CreateTable()
	_, err3 := s.Insert(u1, u2)
	if err1 != nil || err2 != nil || err3 != nil {
		t.Error("failed to init records")
	}

	return s
}
