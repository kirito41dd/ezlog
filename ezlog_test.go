/*
 * Author: zsh
 * Create: 2020/2/14
 */

package ezlog

import (
	"os"
	"testing"
)


func Test_logger(t *testing.T) {
	log := New(os.Stdout, "", 0, LogAll)
	log.SetLogLevel(LogAll)
	log.Debug("debug")
	log.Debugf("%s","debugf")
	log.Info("info")
	log.Infof("%s", "infof")
	log.Warn("warn")
	log.Warnf("%s", "Warnf")
	log.Error("error")
	log.Errorf("%s", "Errorf")
	//log.Panic("panic")
	//log.Panicf("%s", "Panicf")
	//log.Fatal("fatal")
	//log.Fatalf("%s", "Fatalf")
}

func Test_std(t *testing.T) {
	SetLogLevel(LogAll)
	SetFlags(Flags()|BitShortFile)
	Debug("test Debug")
	Debugf("%s","test Debugf")
	Info("test Info")
	Infof("%s", "test Infof")
	Warn("test Warn")
	Warnf("%s", "test Warnf")
	Error("test Error")
	Errorf("%s", "test Errorf")
}