package log

import (
	"github.com/blueskyfish/server-users/configuration"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
	"io"
)

type HttpLogger struct {
	log *logrus.Entry
}

func NewHttpLogger(conf *configuration.Configuration, group string) *HttpLogger {
	return &HttpLogger{
		log: conf.WithGroup(group),
	}
}

func (l HttpLogger) Level() log.Lvl {
	return log.OFF
}

func (l HttpLogger) Prefix() string {
	return ""
}

func (l HttpLogger) SetPrefix(prefix string) {
	// ignore
}

func (l HttpLogger) SetLevel(v log.Lvl) {
	// ignore
}

func (l HttpLogger) SetHeader(h string) {
	// ignore
}

func (l HttpLogger) Print(i ...interface{}) {
	l.log.Print(i...)
}

func (l HttpLogger) Printf(format string, args ...interface{}) {
	l.log.Printf(format, args...)
}

func (l HttpLogger) Printj(fields log.JSON) {
	l.log.WithFields(logrus.Fields(fields)).Print()
}

func (l HttpLogger) Debug(i ...interface{}) {
	l.log.Debug(i...)
}

func (l HttpLogger) Debugf(format string, args ...interface{}) {
	l.log.Debugf(format, args...)
}

func (l HttpLogger) Debugj(fields log.JSON) {
	l.log.WithFields(logrus.Fields(fields)).Debug()
}

func (l HttpLogger) Info(i ...interface{}) {
	l.log.Info(i...)
}

func (l HttpLogger) Infof(format string, args ...interface{}) {
	l.log.Infof(format, args...)
}

func (l HttpLogger) Infoj(fields log.JSON) {
	l.log.WithFields(logrus.Fields(fields)).Info()
}

func (l HttpLogger) Warn(args ...interface{}) {
	l.log.Warn(args...)
}

func (l HttpLogger) Warnf(format string, args ...interface{}) {
	l.log.Warnf(format, args...)
}

func (l HttpLogger) Warnj(fields log.JSON) {
	l.log.WithFields(logrus.Fields(fields)).Warn()
}

func (l HttpLogger) Error(args ...interface{}) {
	l.log.Errorln(args...)
}

func (l HttpLogger) Errorf(format string, args ...interface{}) {
	l.log.Errorf(format, args...)
}

func (l HttpLogger) Errorj(fields log.JSON) {
	l.log.WithFields(logrus.Fields(fields)).Error()
}

func (l HttpLogger) Fatal(args ...interface{}) {
	l.log.Fatal(args...)
}

func (l HttpLogger) Fatalj(fields log.JSON) {
	l.log.WithFields(logrus.Fields(fields)).Fatal()
}

func (l HttpLogger) Fatalf(format string, args ...interface{}) {
	l.log.Fatalf(format, args...)
}

func (l HttpLogger) Panic(args ...interface{}) {
	l.log.Panicln(args...)
}

func (l HttpLogger) Panicj(fields log.JSON) {
	l.log.WithFields(logrus.Fields(fields)).Panic()
}

func (l HttpLogger) Panicf(format string, args ...interface{}) {
	l.log.Panicf(format, args...)
}

func (l HttpLogger) Output() io.Writer {
	return l.log.Logger.Out
}

func (l HttpLogger) SetOutput(w io.Writer) {
	l.log.Logger.SetOutput(w)
}
