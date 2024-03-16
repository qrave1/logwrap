package logwrap

import "context"

type Logger interface {
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
	Info(v ...interface{})
	Infof(format string, v ...interface{})
	Warn(v ...interface{})
	Warnf(format string, v ...interface{})
	Error(v ...interface{})
	Errorf(format string, v ...interface{})
}

// WithContext Создаёт контекст с логгером
func WithContext(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, "ctxLoggerKey", logger)
}

// FromContext Достаёт контекст из логгера
// Иначе возвращает значения из параметра
func FromContext(ctx context.Context, logger Logger) Logger {
	if log, ok := ctx.Value("ctxLoggerKey").(Logger); ok {
		return log
	}
	return logger
}
