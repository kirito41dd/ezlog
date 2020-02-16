/*
 * Author: zsh
 * Create: 2020/2/14
 */

package ezlog

import (
	"fmt"
	"io"
	"os"
)

var std = New(os.Stderr, "", BitStdFlag, LogAll)

func Debug(v ...interface{}) {
	_ = std.Output(2, LogDebug, fmt.Sprintln(v...))
}

func Debugf(format string, v ...interface{}) {
	_ = std.Output(2, LogDebug, fmt.Sprintf(format, v...))
}

func Info(v ...interface{}) {
	_ = std.Output(2, LogInfo, fmt.Sprintln(v...))
}

func Infof(format string, v ...interface{}) {
	_ = std.Output(2, LogInfo, fmt.Sprintf(format, v...))
}

func Warn(v ...interface{}) {
	_ = std.Output(2, LogWarn, fmt.Sprintln(v...))
}

func Warnf(format string, v ...interface{}) {
	_ = std.Output(2, LogWarn, fmt.Sprintf(format, v...))
}

func Error(v ...interface{}) {
	_ = std.Output(2, LogError, fmt.Sprintln(v...))
}

func Errorf(format string, v ...interface{}) {
	_ = std.Output(2, LogError, fmt.Sprintf(format, v...))
}

func Panic(v ...interface{}) {
	s := fmt.Sprintln(v...)
	_ = std.Output(2, LogPanic, s)
	panic(s)
}

func Panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	_ = std.Output(2, LogPanic, s)
	panic(s)
}

func Fatal(v ...interface{}) {
	_ = std.Output(2, LogFatal, fmt.Sprintln(v...))
	os.Exit(1)
}

func Fatalf(format string, v ...interface{}) {
	_ = std.Output(2, LogFatal, fmt.Sprintf(format, v...))
	os.Exit(1)
}


// get and set

func SetOutput(writer io.Writer) {
	std.SetOutput(writer)
}

func Flags() int {
	return std.Flags()
}

func SetFlags(flag int) {
	std.SetFlags(flag)
}

func Prefix() string {
	return std.Prefix()
}

func SetPrefix(s string) {
	std.SetPrefix(s)
}

func LogLevel() int {
	return std.LogLevel()
}

func SetLogLevel(level int) {
	std.SetLogLevel(level)
}

