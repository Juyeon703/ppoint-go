package logue

import (
	"fmt"
	"log"
)

const (
	defaultDepth = 2
)

var (
	callDepth int = 2
)

type Logue struct {
	level Level

	fatal   *log.Logger
	error   *log.Logger
	warning *log.Logger
	info    *log.Logger
	debug   *log.Logger
}

// 로그 인스턴스 생성
func New(level Level, depth int, option LogueOptioner) (*Logue, error) {
	l := &Logue{}
	l.level = level

	callDepth = depth
	if depth < 0 {
		callDepth = defaultDepth
	}
	w, err := option.Setup()

	// Logger 생성
	switch l.level {
	case DEBUG:
		l.debug = log.New(w, "[DEBUG] ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)
		fallthrough
	case INFO:
		l.info = log.New(w, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)
		fallthrough
	case WARNING:
		l.warning = log.New(w, "[WARN] ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)
		fallthrough
	case ERROR:
		l.error = log.New(w, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)
		fallthrough
	case FATAL:
		l.fatal = log.New(w, "[FATAL] ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)
	default:
		return nil, fmt.Errorf("The value of the level variable is not valid.")
	}

	return l, err
}

func (l *Logue) Fatal(v ...interface{}) {
	if l.level < FATAL {
		return
	}

	l.fatal.Output(callDepth, fmt.Sprintln(v...))
}

func (l *Logue) Error(v ...interface{}) {
	if l.level < ERROR {
		return
	}

	l.error.Output(callDepth, fmt.Sprintln(v...))
}

func (l *Logue) Warn(v ...interface{}) {
	if l.level < WARNING {
		return
	}

	l.warning.Output(callDepth, fmt.Sprintln(v...))
}

func (l *Logue) Info(v ...interface{}) {
	if l.level < INFO {
		return
	}

	l.info.Output(callDepth, fmt.Sprintln(v...))
}

func (l *Logue) Debug(v ...interface{}) {
	if l.level < DEBUG {
		return
	}

	v = append(v, []interface{}{trace()})
	l.debug.Output(callDepth, fmt.Sprintln(v...))
}

func (l *Logue) Fatalf(format string, v ...interface{}) {
	if l.level < FATAL {
		return
	}

	l.fatal.Output(callDepth, fmt.Sprintf(format, v...))
}

func (l *Logue) Errorf(format string, v ...interface{}) {
	if l.level < ERROR {
		return
	}

	l.error.Output(callDepth, fmt.Sprintf(format, v...))
}

func (l *Logue) Warnf(format string, v ...interface{}) {
	if l.level < WARNING {
		return
	}

	l.warning.Output(callDepth, fmt.Sprintf(format, v...))
}

func (l *Logue) Infof(format string, v ...interface{}) {
	if l.level < INFO {
		return
	}

	l.info.Output(callDepth, fmt.Sprintf(format, v...))
}

func (l *Logue) Debugf(format string, v ...interface{}) {
	if l.level < DEBUG {
		return
	}

	v = append(v, []interface{}{trace()})
	format = format + " %v"
	l.debug.Output(callDepth, fmt.Sprintf(format, v...))
}
