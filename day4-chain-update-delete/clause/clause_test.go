package clause

import (
	"reflect"
	"testing"
)

func TestClause_Build(t *testing.T) {
	t.Run("select", func(t *testing.T) {
		testSelect(t)
	})
}

func testSelect(t *testing.T)  {
	var clause Clause
	clause.Set(LIMIT, 3)
	clause.Set(WHERE, "Name = ?", "Tom")
	clause.Set(SELECT, "User", []string{"*"})
	clause.Set(ORDERBY, "Age ASC")
	sql, vars := clause.Build(SELECT, WHERE, ORDERBY, LIMIT)
	t.Log(sql, vars)
	want := "SELECT * FROM User WHERE Name = ? ORDER BY Age ASC LIMIT ?"
	if sql != want {
		t.Errorf("want sql \"%s\" but got \"%s\"", want, sql)
	}
	if !reflect.DeepEqual(vars, []interface{}{"Tom", 3}) {
		t.Error("failed to build sql vars")
	}
}