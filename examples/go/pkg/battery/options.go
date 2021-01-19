package battery

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

// Option is the list of options
type Option func(*options)

type options struct {
	target string
	token  string
}

// WithViper loads a config from environment variables.
func WithViper(args ...string) Option {
	return func(o *options) {
		if err := o.viperize(args...); err != nil {
			log.Fatal(err)
		}
	}
}

func (o *options) viperize(args ...string) error {
	v := viper.New()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	v.SetEnvPrefix("ESCAPE_POD_ROBOT")
	v.AutomaticEnv()

	if x := "target"; v.IsSet(x) {
		o.target = v.GetString(x)
	}

	if x := "token"; v.IsSet(x) {
		o.token = v.GetString(x)
	}

	return nil
}
