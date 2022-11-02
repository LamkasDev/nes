package main

import (
	"fmt"
	"runtime"
	"strings"
)

var LogInfoEnabled = true
var LogTraceEnabled = true
var LogWarnEnabled = true

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func LogLn(format string, v ...any) {
	if !LogInfoEnabled {
		return
	}
	Print("[info]", 2, format+"\n", v...)
}

func LogWarnLn(format string, v ...any) {
	if !LogWarnEnabled {
		return
	}
	Print("[warn]", 2, format+"\n", v...)
}

func Print(level string, depth int, format string, v ...any) {
	_, file, line, ok := runtime.Caller(depth)
	if !ok {
		file = "???"
		line = 0
	}
	file = file[(strings.LastIndex(file, "/") + 1):]

	fmt.Printf("[%s] %s: %s", fmt.Sprintf("%s:%v", file, line), level, fmt.Sprintf(format, v...))
}
