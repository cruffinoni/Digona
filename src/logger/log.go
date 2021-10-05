package logger

import (
	"fmt"
	"log"
	"time"
)

type Logger struct{}

func (Logger) Logf(msg string, format ...interface{}) {
	msg = time.Now().Format(time.RFC3339) + " " + msg
	fmt.Printf(msg+ColorReset, format...)
}

func (Logger) Log(msg ...interface{}) {
	fmt.Print(time.Now().Format(time.RFC3339) + " ")
	fmt.Print(msg...)
	fmt.Print(ColorReset)
}

func (Logger) Error(msg ...interface{}) {
	log.Print(msg...)
}

func (Logger) Errorf(msg string, fmt ...interface{}) {
	log.Printf(msg, fmt...)
}

func (Logger) Fatal() {
	log.Panic()
}
func (Logger) FatalMsg(fmt ...interface{}) {
	log.Fatal(fmt...)
}

func (Logger) Fatalf(msg string, fmt ...interface{}) {
	log.Fatalf(msg, fmt...)
}
