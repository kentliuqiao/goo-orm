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

func TestSession_DeleteAndCount(t *testing.T) {
	s := testRecordInit(t)
	affected, err := s.Where("Name = ?", "Tom").Delete()
	count, err := s.Count()

	if err != nil || affected != 1 || count != 1 {
		t.Error("failed to delete or count")
	}
}

func TestSession_Update(t *testing.T) {
	s := testRecordInit(t)
	affected, err := s.Where("Name = ?", "Tom").Update("Age", 100)
	u := &User{}
	err = s.OrderBy("Age DESC").First(u)
	if err != nil || affected != 1 || u.Name != "Tom" {
		t.Error("failed to update")
	}
}

func TestSession_Limit(t *testing.T) {
	s := testRecordInit(t)
	var users []User
	err := s.Limit(1).Find(&users)
	if err != nil || len(users) != 1 {
		t.Error("failed to limit")
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
