package wire

import (
	"io"
)

type Message struct {
}

type MySQLConnectionString string
type Options struct {
	// Messages is the set of recommended greetings.
	Messages []Message
	// Writer is the location to send greetings. nil goes to stdout.
	Writer io.Writer
}

/*func NewGreeter(ctx context.Context, opts *Options) (*Greeter, error) {
	// ...
}

var GreeterSet = wire.NewSet(wire.Struct(new(Options), "*"), NewGreeter)*/
