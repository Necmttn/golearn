package trace

import (
	"fmt"
	"io"
)

// tracer is the interface that describes an Object capable of tracing
// events throughout code.

type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

type nilTracer struct{}

func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}

func (t *nilTracer) Trace(a ...interface{}) {}

// creates new tracer
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

// creates nil tracer
func Off() Tracer {
	return &nilTracer{}
}
