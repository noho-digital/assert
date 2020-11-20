package assert

import (
	"github.com/noho-digital/logging"
)

func NewLogger(l logging.Logger) Logger {
	return log{
		logger: l,
		debug:  New(logHandler{l.Debugf}),
		info:   New(logHandler{l.Infof}),
		panic:  New(logHandler{l.Panicf}),
		error:  New(logHandler{l.Errorf}),
		warn:   New(logHandler{l.Warnf}),
		fatal:  New(logHandler{l.Fatalf}),
	}
}

type Logger interface {
	Debug() *Assertions
	Info() *Assertions
	Warn() *Assertions
	Error() *Assertions
	Panic() *Assertions
	Fatal() *Assertions
}

type logHandler struct {
	fn func(string, ...interface{})
}

func (l logHandler) Errorf(format string, args ...interface{}) {
	l.fn(format, args...)
}

type log struct {
	logger logging.Logger
	info   *Assertions
	panic  *Assertions
	error  *Assertions
	warn   *Assertions
	debug  *Assertions
	fatal  *Assertions
}

func (l log) Debug() *Assertions {
	return l.debug
}

func (l log) Info() *Assertions {
	return l.info
}

func (l log) Warn() *Assertions {
	return l.warn
}

func (l log) Error() *Assertions {
	return l.error
}

func (l log) Panic() *Assertions {
	return l.panic
}

func (l log) Fatal() *Assertions {
	return l.fatal
}
