// Package trace provides helper functions that let developers easily log the
// execution flow of their programs.
package trace

import (
	"log"
	"os"
	"runtime"
)

// Line returns the function name and line number wherever Line() is
// called. This is most useful when creating a trail of call executions across
// multiple goroutines.
func Line() {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	log.Printf("%s:%d\n", frame.Function, frame.Line)
}

// Error returns the function name, line number wherever Error() is called, and
// the supplied error message. This is most useful when creating a trail of
// errors across multiple goroutines.
func Error(err error) {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	log.Printf("%s:%d: %s\n", frame.Function, frame.Line, err.Error())
}

// DebugLine is the same as Line(), except that it only prints out information
// when the DEBUG environment variable is set to "true".
func DebugLine() {
	debug := os.Getenv("DEBUG")
	if debug != "true" {
		return
	}

	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	log.Printf("%s:%d\n", frame.Function, frame.Line)
}

// DebugMessage prints the supplied message only when the DEBUG environment
// variable is set to "true".
func DebugMessage(msg string, v ...interface{}) {
	debug := os.Getenv("DEBUG")
	if debug != "true" {
		return
	}

	log.Printf(msg, v...)
}
