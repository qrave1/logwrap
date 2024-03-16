package logrus

import (
	"github.com/sirupsen/logrus"
	"os"
)

type LogrusWrapper struct {
	*logrus.Entry
}

func NewWrapper(level logrus.Level, formatter logrus.Formatter) *LogrusWrapper {
	lr := logrus.New()
	lr.SetLevel(level)
	lr.Formatter = formatter
	lr.SetOutput(os.Stdout)
	lr.SetNoLock()

	log := &LogrusWrapper{Entry: logrus.NewEntry(lr)}

	return log
}

func NewDefaultTextWrapper() *LogrusWrapper {
	lr := logrus.New()
	lr.SetLevel(logrus.InfoLevel)
	lr.Formatter = &logrus.TextFormatter{}
	lr.SetOutput(os.Stdout)
	lr.SetNoLock()

	log := &LogrusWrapper{Entry: logrus.NewEntry(lr)}

	return log
}

func NewDefaultPrettyWrapper() *LogrusWrapper {
	lr := logrus.New()
	lr.SetLevel(logrus.InfoLevel)
	lr.Formatter = &PrettyFormatter{}
	lr.SetOutput(os.Stdout)
	lr.SetNoLock()

	log := &LogrusWrapper{Entry: logrus.NewEntry(lr)}

	return log
}
