/*
 * Author: zsh
 * Create: 2020/2/14
 */

package ezlog

import (
	"io"
	"os"
)

var std = New(os.Stderr, "", BitStdFlag, LogAll)

func Debug(v ...interface{}) {
	std.Debug(v...)
}

func Debugf(format string, v ...interface{}) {
	std.Debugf(format, v...)
}

func Info(v ...interface{}) {
	std.Info(v...)
}

func Infof(format string, v ...interface{}) {
	std.Infof(format, v...)
}

func Warn(v ...interface{}) {
	std.Warn(v...)
}

func Warnf(format string, v ...interface{}) {
	std.Warnf(format, v...)
}

func Error(v ...interface{}) {
	std.Error(v...)
}

func Errorf(format string, v ...interface{}) {
	std.Errorf(format, v...)
}

func Panic(v ...interface{}) {
	std.Panic(v...)
}

func Panicf(format string, v ...interface{}) {
	std.Panicf(format, v...)
}

func Fatal(v ...interface{}) {
	std.Fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
	std.Fatalf(format, v...)
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

