package gooorm

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestNewEngine(t *testing.T) {
	engine := OpenDB(t)
	defer engine.Close()
}

func OpenDB(t *testing.T) *Engine {
	t.Helper()
	engine, err := NewEngine("sqlite3", "goo.db")
	if err != nil {
		t.Error("failed to connect", err)
	}

	return engine
}
