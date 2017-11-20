package trace

import (
	"io"
)

// tracer is the interface that describes an Object capable of tracing
// events throughout code.

type Tracer interface {
	Trace(...interface{})
}

func New(w io.Writer) {}
