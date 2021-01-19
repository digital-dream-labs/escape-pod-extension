package server

import (
	"context"
	"fmt"
)

type Processor interface {
	Get(context.Context, string, map[string]string) (string, map[string]string, error)
}

// Server is the configuration struct
type Server struct {
	processors map[string]Processor
}

// New returns a populated Router struct
func New(opts ...Option) (*Server, error) {
	cfg := options{}

	for _, opt := range opts {
		opt(&cfg)
	}
	if cfg.processors == nil {
		return nil, fmt.Errorf("no processors defined")
	}

	r := Server{
		processors: cfg.processors,
	}
	return &r, nil
}
