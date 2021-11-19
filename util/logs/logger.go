package logs

import "log"

var logs log.Logger

func Info(message interface{}) {
	logs.Printf("info: %v", message)
}

func Error(message interface{}, payload interface{}) {
	logs.Printf("error: %v, data: %v", message, payload)
}

func Panic(message interface{}, payload interface{}) {
	logs.Panicf("error: %v, data: %v", message, payload)
}

func Fatal(message interface{}, payload interface{}) {
	logs.Fatalf("error: %v, data: %v", message, payload)
}

func init() {
	logs = *log.Default()
}
