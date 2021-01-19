package genericresponse

import (
	"fmt"

	sdk "github.com/digital-dream-labs/vector-go-sdk/pkg/vector"
)

// GenericResponse is the configuration struct
type GenericResponse struct {
	bot      *sdk.Vector
	response string
}

// New returns a populated IntentHack struct
func New(opts ...Option) (*GenericResponse, error) {
	cfg := options{}

	for _, opt := range opts {
		opt(&cfg)
	}

	if cfg.response == "" {
		return nil, fmt.Errorf("please define a response")
	}

	bot, err := sdk.New(
		sdk.WithToken(cfg.token),
		sdk.WithTarget(cfg.target),
	)
	if err != nil {
		return nil, err
	}

	wu := GenericResponse{
		bot:      bot,
		response: cfg.response,
	}

	return &wu, nil
}
