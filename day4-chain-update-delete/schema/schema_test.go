package schema

import (
	"gooorm/dialect"
	"reflect"
	"testing"
)

type User struct {
	Name string `gooorm:"PRIMARY KEY"`
	Age int
	Address string
}

var TestDialect, _ = dialect.GetDialect("sqlite3")

func TestParse(t *testing.T) {
	schema := Parse(&User{}, TestDialect)
	if schema.Name != "User" || len(schema.Fields) != 2 {
		t.Error("failed to parse User struct")
	}
	if schema.GetField("Name").Tag != "PRIMARY KEY" {
		t.Error("failed to parse tag")
	}
}

func TestSchema_RecordValues(t *testing.T) {
	schema := Parse(&User{}, TestDialect)
	u1 := &User{
		Name: "Tom",
		Age: 12,
	}
	u2 := &User{
		Name: "Jerry",
		Age: 14,
	}
	f1 := schema.RecordValues(u1)
	want := []interface{}{"Tom", 12, ""}
	if !reflect.DeepEqual(f1, want) {
		t.Errorf("got %+v, but want %+v", f1, want)
	}

	f2 := schema.RecordValues(u2)
	want = []interface{}{"Jerry", 14, ""}
	if !reflect.DeepEqual(f2, want) {
		t.Errorf("got %+v, but want %+v", f2, want)
	}
}

func TestReflect_NumField(t *testing.T) {
	type Book struct {
		Name string
		Author string
		Year int
	}
	b1 := &Book{
		Name: "1984",
		Year: 1984,
	}
	modelTyp := reflect.Indirect(reflect.ValueOf(b1)).Type()
	numField := modelTyp.NumField()
	if numField != 3 {
		t.Errorf("got %d but want %d", numField, 3)
	}
}
