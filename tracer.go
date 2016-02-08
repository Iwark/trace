package trace

import (
	"fmt"
	"io"
)

// Tracer is an object capable of
// tracing events throughout code.
type Tracer struct {
	out io.Writer
}

// New creates a new Tracer that will write the output to
// the specified io.Writer.
func New(w io.Writer) Tracer {
	return Tracer{out: w}
}

// Trace writes the arguments to this Tracers io.Writer.
func (t Tracer) Trace(a ...interface{}) {
	if t.out == nil {
		return
	}
	fmt.Fprintln(t.out, a...)
}
