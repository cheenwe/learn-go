package logger

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

const (
	logInfoLabel  = "[info] "
	logTraceLabel = "[trace] "
	logWarnLabel  = "[warn] "
	logErrorLabel = "[error] "
)

type Context interface{}

type cidContext interface {
	Cid() int
}

type loggerPlus struct {
	logger *log.Logger
}

func NewLoggerPlus(l *log.Logger) Logger {
	return &loggerPlus{logger: l}
}

func (v *loggerPlus) format(ctx Context, a ...interface{}) []interface{} {
	if ctx == nil {
		return append([]interface{}{fmt.Sprintf("[%v] ", os.Getpid())}, a...)
	} else if ctx, ok := ctx.(cidContext); ok {
		return append([]interface{}{fmt.Sprintf("[%v][%v] ", os.Getpid(), ctx.Cid())}, a...)
	}
	return a
}

func (v *loggerPlus) formatf(ctx Context, format string, a ...interface{}) (string, []interface{}) {
	if ctx == nil {
		return "[%v] " + format, append([]interface{}{os.Getpid()}, a...)
	} else if ctx, ok := ctx.(cidContext); ok {
		return "[%v][%v] " + format, append([]interface{}{os.Getpid(), ctx.Cid()}, a...)
	}
	return format, a
}

var colorYellow = "\033[33m"
var colorRed = "\033[31m"
var colorBlack = "\033[0m"

func (v *loggerPlus) doPrintln(args ...interface{}) {
	if previousIo == nil {
		if v == Error {
			fmt.Fprintf(os.Stdout, colorRed)
			v.logger.Println(args...)
			fmt.Fprintf(os.Stdout, colorBlack)
		} else if v == Warn {
			fmt.Fprintf(os.Stdout, colorYellow)
			v.logger.Println(args...)
			fmt.Fprintf(os.Stdout, colorBlack)
		} else {
			v.logger.Println(args...)
		}
	} else {
		v.logger.Println(args...)
	}
}

func (v *loggerPlus) doPrintf(format string, args ...interface{}) {
	if previousIo == nil {
		if v == Error {
			fmt.Fprintf(os.Stdout, colorRed)
			v.logger.Printf(format, args...)
			fmt.Fprintf(os.Stdout, colorBlack)
		} else if v == Warn {
			fmt.Fprintf(os.Stdout, colorYellow)
			v.logger.Printf(format, args...)
			fmt.Fprintf(os.Stdout, colorBlack)
		} else {
			v.logger.Printf(format, args...)
		}
	} else {
		v.logger.Printf(format, args...)
	}
}

var Info Logger

func I(ctx Context, a ...interface{}) {
	Info.Println(ctx, a...)
}

func If(ctx Context, format string, a ...interface{}) {
	Info.Printf(ctx, format, a...)
}

var Trace Logger

func T(ctx Context, a ...interface{}) {
	Trace.Println(ctx, a...)
}

func Tf(ctx Context, format string, a ...interface{}) {
	Trace.Printf(ctx, format, a...)
}

var Warn Logger

func W(ctx Context, a ...interface{}) {
	Warn.Println(ctx, a...)
}

func Wf(ctx Context, format string, a ...interface{}) {
	Warn.Printf(ctx, format, a...)
}

var Error Logger

func E(ctx Context, a ...interface{}) {
	Error.Println(ctx, a...)
}

func Ef(ctx Context, format string, a ...interface{}) {
	Error.Printf(ctx, format, a...)
}

type Logger interface {
	Println(ctx Context, a ...interface{})
	Printf(ctx Context, format string, a ...interface{})
}

func init() {
	Info = NewLoggerPlus(log.New(ioutil.Discard, logInfoLabel, log.Ldate|log.Ltime|log.Lmicroseconds))
	Trace = NewLoggerPlus(log.New(os.Stdout, logTraceLabel, log.Ldate|log.Ltime|log.Lmicroseconds))
	Warn = NewLoggerPlus(log.New(os.Stdout, logWarnLabel, log.Ldate|log.Ltime|log.Lmicroseconds))
	Error = NewLoggerPlus(log.New(os.Stdout, logErrorLabel, log.Ldate|log.Ltime|log.Lmicroseconds))
}

func Switch(w io.Writer) {
	Info = NewLoggerPlus(log.New(ioutil.Discard, logInfoLabel, log.Ldate|log.Ltime|log.Lmicroseconds))
	Trace = NewLoggerPlus(log.New(w, logTraceLabel, log.Ldate|log.Ltime|log.Lmicroseconds))
	Warn = NewLoggerPlus(log.New(w, logWarnLabel, log.Ldate|log.Ltime|log.Lmicroseconds))
	Error = NewLoggerPlus(log.New(w, logErrorLabel, log.Ldate|log.Ltime|log.Lmicroseconds))

	if w, ok := w.(io.Closer); ok {
		previousIo = w
	}
}

var previousIo io.Closer

func Close() (err error) {
	Info = NewLoggerPlus(log.New(ioutil.Discard, logInfoLabel, log.Ldate|log.Ltime|log.Lmicroseconds))
	Trace = NewLoggerPlus(log.New(ioutil.Discard, logTraceLabel, log.Ldate|log.Ltime|log.Lmicroseconds))
	Warn = NewLoggerPlus(log.New(ioutil.Discard, logWarnLabel, log.Ldate|log.Ltime|log.Lmicroseconds))
	Error = NewLoggerPlus(log.New(ioutil.Discard, logErrorLabel, log.Ldate|log.Ltime|log.Lmicroseconds))

	if previousIo != nil {
		err = previousIo.Close()
		previousIo = nil
	}

	return
}
