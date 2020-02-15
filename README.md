# ezlog
[![Build Status](https://www.travis-ci.org/zshorz/ezlog.svg?branch=master)](https://www.travis-ci.org/zshorz/ezlog)

A simple log package with log levels.

## usage
Example:
```go
package main

import (
	"fmt"
	"github.com/zshorz/ezlog"
	"os"
)

func main() {
	ezlog.Debug("ezlog Debug")
	ezlog.Info("ezlog Info")
	ezlog.Warn("ezlog Warn")

	fmt.Fprintln(os.Stderr, "----------------------------------")
	ezlog.SetLogLevel(ezlog.LogWarn)
	ezlog.Debug("ezlog Debug")
	ezlog.Info("ezlog Info")
	ezlog.Warn("ezlog Warn")
	ezlog.Error("ezlog Error")

	fmt.Fprintln(os.Stderr, "----------------------------------")
	log := ezlog.New(os.Stderr, "mylog", ezlog.BitDate|ezlog.BitTime|ezlog.BitShortFile, ezlog.LogAll)
	log.Info("info 1", "info 2")
	log.SetFlags(log.Flags()|ezlog.BitMicroSeconds)
	log.Infof("my age is %d", 20)

	fmt.Fprintln(os.Stderr, "----------------------------------")
	log.SetLogLevel(ezlog.LogOff)
	log.Error("error")
}
```
Output:
```text
2020/02/14 21:04:16 [DEBUG] ezlog Debug
2020/02/14 21:04:16 [INFO ] ezlog Info
2020/02/14 21:04:16 [WARN ] ezlog Warn
----------------------------------
2020/02/14 21:04:16 [WARN ] ezlog Warn
2020/02/14 21:04:16 [ERROR] ezlog Error
----------------------------------
<mylog> 2020/02/14 21:04:16 [INFO ] main.go:23: info 1 info 2
<mylog> 2020/02/14 21:04:16.998167 [INFO ] main.go:25: my age is 20
----------------------------------

(Process finished with exit code 0)
```

## go doc -all
```godoc
package ezlog // import "github.com/zshorz/ezlog"


CONSTANTS

const (
        BitDate = 1 << iota
        BitTime
        BitMicroSeconds
        BitLongFile
        BitShortFile
        BitUTC                         // use UTC rather than the local time zone
        BitStdFlag = BitDate | BitTime // initial values for the standard logger
        BitDefault = BitStdFlag | BitShortFile
)
    Flag

const (
        LogDebug = iota
        LogInfo
        LogWarn
        LogError
        LogPanic
        LogFatal
        LogOff

        LogAll = LogDebug
)
    Debug level


FUNCTIONS

func Debug(v ...interface{})
func Debugf(format string, v ...interface{})
func Error(v ...interface{})
func Errorf(format string, v ...interface{})
func Fatal(v ...interface{})
func Fatalf(format string, v ...interface{})
func Flags() int
func Info(v ...interface{})
func Infof(format string, v ...interface{})
func LogLevel() int
func Panic(v ...interface{})
func Panicf(format string, v ...interface{})
func Prefix() string
func SetFlags(flag int)
func SetLogLevel(level int)
func SetPrefix(s string)
func Version() string
    return version as string

func Warn(v ...interface{})
func Warnf(format string, v ...interface{})

TYPES

type EzLogger struct {
        // Has unexported fields.
}

func New(out io.Writer, prefix string, flag int, level int) *EzLogger
    New creates a new EzLogger. The out variable sets the destination to which
    log data will be written. The prefix appears at the beginning of each
    generated log line. The flag argument defines the logging properties. The
    level set the lowest level of log output

func (ezlog *EzLogger) Debug(v ...interface{})

func (ezlog *EzLogger) Debugf(format string, v ...interface{})

func (ezlog *EzLogger) Error(v ...interface{})

func (ezlog *EzLogger) Errorf(format string, v ...interface{})

func (ezlog *EzLogger) Fatal(v ...interface{})

func (ezlog *EzLogger) Fatalf(format string, v ...interface{})

func (ezlog *EzLogger) Flags() int

func (ezlog *EzLogger) Info(v ...interface{})

func (ezlog *EzLogger) Infof(format string, v ...interface{})

func (ezlog *EzLogger) LogLevel() int

func (ezlog *EzLogger) Output(callDepth int, level int, s string) error
    Output writes the output for a logging event. callDepth is used to recover
    the PC and is provided for generality, on all pre-defined it will be 2.
    level is log level s is the data

func (ezlog *EzLogger) Panic(v ...interface{})

func (ezlog *EzLogger) Panicf(format string, v ...interface{})

func (ezlog *EzLogger) Prefix() string

func (ezlog *EzLogger) SetFlags(flag int)

func (ezlog *EzLogger) SetLogLevel(level int)

func (ezlog *EzLogger) SetOutput(w io.Writer)
    SetOutput resets the output destination for the logger.

func (ezlog *EzLogger) SetPrefix(s string)

func (ezlog *EzLogger) Warn(v ...interface{})

func (ezlog *EzLogger) Warnf(format string, v ...interface{})
```
