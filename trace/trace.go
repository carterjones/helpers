// Package trace provides helper functions that let developers easily log the
// execution flow of their programs.
package trace

import (
	"log"
	"os"
	"runtime"
	"strings"
)

func traceEnabled() bool {
	v := os.Getenv("TRACE")
	return strings.ToLower(v) == "true"
}

// Line returns the function name and line number wherever Line() is
// called. This is most useful when creating a trail of call executions across
// multiple goroutines.
func Line() {
	if !traceEnabled() {
		return
	}

	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	log.Printf("%s:%d\n", frame.Function, frame.Line)
}
