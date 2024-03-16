package logrus

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
	"testing"
	"time"
)

func TestPrettyFormatter_Format(t *testing.T) {
	f := PrettyFormatter{}

	tests := []struct {
		name string
		log  *logrus.Entry
	}{
		{
			name: "Debug",
			log:  newLog(logrus.DebugLevel),
		},
		{
			name: "Info",
			log:  newLog(logrus.InfoLevel),
		},
		{
			name: "Warn",
			log:  newLog(logrus.WarnLevel),
		},
		{
			name: "Error",
			log:  newLog(logrus.ErrorLevel),
		},
	}
	for _, tt := range tests {
		b, _ := f.Format(tt.log)

		expected := strings.Join([]string{
			fmt.Sprintf("[%s\x1b[0m]", lvlFromLog(tt.log.Level)),
			tt.log.Time.Format(defaultTimestampFormat),
			"- Test Message\n",
		}, " ")

		if string(b) != expected {
			t.Errorf("formatting expected result was %q instead of %q", string(b), expected)
		}
	}
}

func newLog(level logrus.Level) *logrus.Entry {
	e := logrus.WithField("", "")
	e.Message = "Test Message"
	e.Level = level
	e.Time = time.Now()

	return e
}

func lvlFromLog(lvl logrus.Level) string {
	switch lvl {
	case logrus.DebugLevel:
		return "\u001B[0;36mDEBUG"
	case logrus.InfoLevel:
		return "\u001B[0;32mINFO"
	case logrus.WarnLevel:
		return "\u001B[0;33mWARN"
	case logrus.ErrorLevel:
		return "\u001B[0;31mERROR"
	default:
	}

	return ""
}
