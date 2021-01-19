package server

type options struct {
	processors map[string]Processor
}

// Option is the list of options
type Option func(*options)

// WithProcessors sets the intent processor
func WithProcessors(s map[string]Processor) Option {
	return func(o *options) {
		o.processors = s
	}
}
