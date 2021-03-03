package log

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
)

// log levels
const (
	InfoLevel = iota
	ErrorLevel
	Disabled
)

var (
	errorLog = log.New(os.Stdout, "\033[31m[ERROR]\033[0m ", log.LstdFlags|log.Lshortfile)
	infoLog  = log.New(os.Stdout, "\033[34m[INFO ]\033[0m ", log.LstdFlags|log.Lshortfile)
	loggers  = []*log.Logger{errorLog, infoLog}
	mu       sync.Mutex
)

// log methods
var (
	Error  = errorLog.Println
	Errorf = errorLog.Printf
	Info   = infoLog.Println
	Infof  = infoLog.Printf
)

// SetLevel controls log level
func SetLevel(level int) {
	mu.Lock()
	defer mu.Unlock()

	for _, logger := range loggers {
		logger.SetOutput(os.Stdout)
	}

	if ErrorLevel < level {
		errorLog.SetOutput(ioutil.Discard)
	}
	if InfoLevel < level {
		infoLog.SetOutput(ioutil.Discard)
	}
}
