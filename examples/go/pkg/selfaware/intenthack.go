package selfaware

import sdk "github.com/digital-dream-labs/vector-go-sdk/pkg/vector"

// SelfAware is the configuration struct
type SelfAware struct {
	bot *sdk.Vector
}

// New returns a populated SelfAware struct
func New(opts ...Option) (*SelfAware, error) {
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

	wu := SelfAware{
		bot: bot,
	}

	return &wu, nil
}
