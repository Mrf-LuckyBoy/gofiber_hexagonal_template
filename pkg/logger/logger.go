package logger

import (
	"log"
	"os"
)

var L = log.New(os.Stdout, "[app] ", log.LstdFlags|log.Lshortfile)
