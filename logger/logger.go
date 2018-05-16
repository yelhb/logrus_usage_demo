package logger

import (
	"io"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func SetOutput(w io.Writer) {
	log.Out = w
}

func SetFormatter(f logrus.Formatter) {
	log.Formatter = f
}

func SetLevel(l logrus.Level) {
	log.Level = l
}

func Warning(args ...interface{}) {
	log.Warning(args)
}

func Info(args ...interface{}) {
	log.Info(args)
}

func Debug(args ...interface{}) {
	log.Debug(args)
}
