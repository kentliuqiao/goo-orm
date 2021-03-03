package schema

import (
	"gooorm/dialect"
	"testing"
)

type User struct {
	Name string `gooorm:"PRIMARY KEY"`
	Age int
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
