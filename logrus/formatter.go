package logrus

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

const (
	defaultTimestampFormat = "2006-01-02 15:04:05"
)

type PrettyFormatter struct {
	TimestampFormat string
}

func (p *PrettyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestampFormat := p.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}

	log := fmt.Sprintf(
		"[%s] %s - %s\n",
		color(entry.Level),
		entry.Time.Format(timestampFormat),
		entry.Message,
	)

	return []byte(log), nil
}

func color(lvl logrus.Level) string {
	switch lvl {
	case logrus.DebugLevel:
		return fmt.Sprintf("\u001B[0;36m%s\u001B[0m", "DEBUG") // cyan color
	case logrus.InfoLevel:
		return fmt.Sprintf("\u001B[0;32m%s\u001B[0m", "INFO") // green color
	case logrus.WarnLevel:
		return fmt.Sprintf("\u001B[0;33m%s\u001B[0m", "WARN") // yellow color
	case logrus.ErrorLevel:
		return fmt.Sprintf("\u001B[0;31m%s\u001B[0m", "ERROR") // red color
	default:
	}

	return ""
}
