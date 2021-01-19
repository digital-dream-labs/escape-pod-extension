package battery

import sdk "github.com/digital-dream-labs/vector-go-sdk/pkg/vector"

// IntentHack is the configuration struct
type Battery struct {
	bot *sdk.Vector
}

// New returns a populated IntentHack struct
func New(opts ...Option) (*Battery, error) {
	cfg := options{}

	for _, opt := range opts {
		opt(&cfg)
	}

	bot, err := sdk.New(
		sdk.WithToken(cfg.token),
		sdk.WithTarget(cfg.target),
	)
	if err != nil {
		return nil, err
	}

	wu := Battery{
		bot: bot,
	}

	return &wu, nil
}
