/*
 * Author: zsh
 * Create: 2020/2/14
 */

package ezlog

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
	"time"
)

// Flag
const (
	BitDate		= 1 << iota
	BitTime
	BitMicroSeconds
	BitLongFile
	BitShortFile
	BitUTC			// use UTC rather than the local time zone
	BitStdFlag	= BitDate | BitTime // initial values for the standard logger
	BitDefault	= BitStdFlag | BitShortFile
)

// Debug level
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

// Debug Level string
var levels = []string{
	"[DEBUG]",
	"[INFO ]",
	"[WARN ]",
	"[ERROR]",
	"[PANIC]",
	"[FATAL]",
	"", // Placeholder for 	LogOff
}

type EzLogger struct {
	mutex 	  sync.Mutex		// ensures atomic writes; protects the following fields
	prefix	  string 			// prefix to write at beginning of each line
	flag	  int				// flag properties
	out		  io.Writer 		// destination for output
	buf 	  bytes.Buffer 		// for accumulating text to write
	level	  int 				// log level
}

// New creates a new EzLogger.
// The out variable sets the destination to which log data will be written.
// The prefix appears at the beginning of each generated log line.
// The flag argument defines the logging properties.
// The level set the lowest level of log output
func New(out io.Writer, prefix string, flag int, level int) *EzLogger {
	return &EzLogger{out: out, prefix: prefix, flag: flag, level: level}
}


// formatHeader writes log header to buf
func (ezlog *EzLogger) formatHeader(buf *bytes.Buffer, t time.Time, file string, line int, level int) {
	// prefix
	if ezlog.prefix != "" {
		buf.WriteByte('<')
		buf.WriteString(ezlog.prefix)
		buf.WriteByte('>')
		buf.WriteByte(' ')
	}

	// date time
	if ezlog.flag & (BitDate|BitTime|BitMicroSeconds) != 0 {
		if ezlog.flag & BitUTC != 0 {
			t = t.UTC()
		}
		// date
		if ezlog.flag & BitDate != 0 {
			year, month, day := t.Date()
			itoa(buf, year, 4)
			buf.WriteByte('/') // 2020/
			itoa(buf, int(month), 2)
			buf.WriteByte('/') // 2020/02/
			itoa(buf, day, 2)
			buf.WriteByte(' ') // 2020/02/14
		}
		// time
		if ezlog.flag & (BitTime|BitMicroSeconds) != 0 {
			hour, min, sec := t.Clock()
			itoa(buf, hour, 2)
			buf.WriteByte(':') // 17:
			itoa(buf, min, 2)
			buf.WriteByte(':') // 17:37:
			itoa(buf, sec, 2)
			if ezlog.flag & BitMicroSeconds != 0 {
				buf.WriteByte('.')
				itoa(buf, t.Nanosecond()/1e3, 6) // 17:37:50.123456
			}
			buf.WriteByte(' ')
		}
	}

	// log level
	buf.WriteString(levels[level])
	buf.WriteByte(' ')

	// file
	if ezlog.flag & (BitShortFile|BitLongFile) != 0 {
		if ezlog.flag & BitShortFile != 0 {
			short := file
			for i := len(file) - 1 ; i > 0; i-- {
				if file[i] == '/' {
					short = file[i+1:]
					break
				}
			}
			file = short
		}
		buf.WriteString(file)
		buf.WriteByte(':')
		// line
		itoa(buf, line, -1)
		buf.WriteString(": ")
	}
}

// Output writes the output for a logging event.
// callDepth is used to recover the PC and is provided for generality, on all pre-defined it will be 2.
// level is log level
// s is the data
func (ezlog *EzLogger) Output(callDepth int, level int, s string) error {
	ezlog.mutex.Lock()
	defer ezlog.mutex.Unlock()

	// check log levele decide whether to output
	if level < ezlog.level {
		return nil
	}

	now := time.Now()
	var file string
	var line int


	if ezlog.flag & (BitShortFile|BitLongFile) != 0 {
		ezlog.mutex.Unlock()
		var ok bool
		_, file, line, ok = runtime.Caller(callDepth)
		if !ok {
			file = "???"
			line = 0
		}
		ezlog.mutex.Lock()
	}
	// clear buf
	ezlog.buf.Reset()
	ezlog.formatHeader(&ezlog.buf, now, file, line, level)
	ezlog.buf.WriteString(s)
	if len(s) > 0 && s[len(s)-1] != '\n' {
		ezlog.buf.WriteByte('\n')
	}
	_, err := ezlog.out.Write(ezlog.buf.Bytes())
	return err
}

// ------ Debug ------

func (ezlog *EzLogger) Debugf(format string, v ...interface{}) {
	_ = ezlog.Output(2, LogDebug, fmt.Sprintf(format, v...))
}

func (ezlog *EzLogger) Debug(v ...interface{}) {
	_ = ezlog.Output(2, LogDebug, fmt.Sprintln(v...))
}

// ------ Info ------

func (ezlog *EzLogger) Infof(format string, v ...interface{}) {
	_ = ezlog.Output(2, LogInfo, fmt.Sprintf(format, v...))
}

func (ezlog *EzLogger) Info(v ...interface{}) {
	_ = ezlog.Output(2, LogInfo, fmt.Sprintln(v...))
}

// ------ Warn ------

func (ezlog *EzLogger) Warnf(format string, v ...interface{}) {
	_ = ezlog.Output(2, LogWarn, fmt.Sprintf(format, v...))
}

func (ezlog *EzLogger) Warn(v ...interface{}) {
	_ = ezlog.Output(2, LogWarn, fmt.Sprintln(v...))
}

// ------ Error ------

func (ezlog *EzLogger) Errorf(format string, v ...interface{}) {
	_ = ezlog.Output(2, LogError, fmt.Sprintf(format, v...))
}

func (ezlog *EzLogger) Error(v ...interface{}) {
	_ = ezlog.Output(2, LogError, fmt.Sprintln(v...))
}

// ------ Panic ------

func (ezlog *EzLogger) Panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	_ = ezlog.Output(2, LogPanic, s)
	panic(s)
}

func (ezlog *EzLogger) Panic(v ...interface{}) {
	s := fmt.Sprintln(v...)
	_ = ezlog.Output(2, LogPanic, s)
	panic(s)
}

// ------ Fatal ------

func (ezlog *EzLogger) Fatalf(format string, v ...interface{}) {
	_ = ezlog.Output(2, LogFatal, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func (ezlog *EzLogger) Fatal(v ...interface{}) {
	_ = ezlog.Output(2, LogFatal, fmt.Sprintln(v...))
	os.Exit(1)
}

// ----- get and set ------

// SetOutput resets the output destination for the logger.
func (ezlog *EzLogger) SetOutput(w io.Writer) {
	ezlog.mutex.Lock()
	defer  ezlog.mutex.Unlock()
	ezlog.out = w
}

func (ezlog *EzLogger) Flags() int {
	ezlog.mutex.Lock()
	defer ezlog.mutex.Unlock()
	return ezlog.flag
}

func (ezlog *EzLogger) SetFlags(flag int) {
	ezlog.mutex.Lock()
	defer ezlog.mutex.Unlock()
	ezlog.flag = flag
}

func (ezlog *EzLogger) Prefix() string {
	ezlog.mutex.Lock()
	defer ezlog.mutex.Unlock()
	return ezlog.prefix
}

func (ezlog *EzLogger) SetPrefix(s string) {
	ezlog.mutex.Lock()
	defer ezlog.mutex.Unlock()
	ezlog.prefix = s
}

func (ezlog *EzLogger) LogLevel() int {
	ezlog.mutex.Lock()
	defer ezlog.mutex.Unlock()
	return ezlog.level
}

func (ezlog *EzLogger) SetLogLevel(level int) {
	ezlog.mutex.Lock()
	defer ezlog.mutex.Unlock()
	ezlog.level = level
}

