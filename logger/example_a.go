package main

import (
	"os"

	ol "github.com/cheenwe/learn-go/logger"
)

func ExampleLogger_ToConsole() {
	// Simply log to console.
	ol.Info.Println(nil, "The log text.")
	ol.Trace.Println(nil, "The log text.")
	ol.Warn.Println(nil, "The log text.")
	ol.Error.Println(nil, "The log text.")

	// Use short aliases.
	ol.I(nil, "The log text.")
	ol.T(nil, "The log text.")
	ol.W(nil, "The log text.")
	ol.E(nil, "The log text.")

	// Use printf style log.
	ol.If(nil, "The log %v", "text")
	ol.Tf(nil, "The log %v", "text")
	ol.Wf(nil, "The log %v", "text")
	ol.Ef(nil, "The log %v", "text")
}

func ExampleLogger_ToFile() {
	// Open logger file and change the tank for logger.
	var err error
	var f *os.File
	if f, err = os.Open("sys.log"); err != nil {
		return
	}
	ol.Switch(f)

	// Use logger, which will write to file.
	ol.T(nil, "The log text.")

	// Close logger file when your application quit.
	defer ol.Close()
}

func ExampleLogger_SwitchFile() {
	// Initialize logger with file.
	var err error
	var f *os.File
	if f, err = os.Open("sys.log"); err != nil {
		return
	}
	ol.Switch(f)

	// When need to reap log file,
	// user must close current log file.
	ol.Close()
	// User can move the sys.log away.
	// Then reopen the log file and notify logger to use it.
	if f, err = os.Open("sys.log"); err != nil {
		return
	}
	// All logs between close and switch are dropped.
	ol.Switch(f)

	// Always close it.
	defer ol.Close()
}

// Each context is specified a connection,
// which user must implement the interface.
type cidContext int

func (v cidContext) Cid() int {
	return int(v)
}

func ExampleLogger_ConnectionBased() {
	ctx := cidContext(100)
	ol.Info.Println(ctx, "The log text")
	ol.Trace.Println(ctx, "The log text.")
	ol.Warn.Println(ctx, "The log text.")
	ol.Error.Println(ctx, "The log text.")
}

func main() {
	ExampleLogger_ToConsole()
	ExampleLogger_ToFile()
	ExampleLogger_SwitchFile()
	ExampleLogger_ConnectionBased()

}
