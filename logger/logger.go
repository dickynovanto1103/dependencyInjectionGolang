package logger

import "fmt"

type Logger struct{}

func (logger *Logger) Log(s string) {
	fmt.Println(s)
}
