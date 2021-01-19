package openweathermap

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

// Option is the list of options
type Option func(*options)

type options struct {
	apiKey     string
	units      string
	locationID string
	cityName   string
	zipcode    string
}

// WithAPIKey sets the API key
func WithAPIKey(s string) Option {
	return func(o *options) {
		o.apiKey = s
	}
}

// WithUnits sets the unit for temperature.  Valid units are metric and imperial.
func WithUnits(s string) Option {
	return func(o *options) {
		o.units = s
	}
}

// WithLocationID sets openweathermap city ID.
func WithLocationID(s string) Option {
	return func(o *options) {
		o.locationID = s
	}
}

// WithCity sets the city name.
func WithCity(s string) Option {
	return func(o *options) {
		o.cityName = s
	}
}

// WithZipCode sets the zip code.
func WithZipCode(s string) Option {
	return func(o *options) {
		o.zipcode = s
	}
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
	v.SetEnvPrefix("ESCAPE_POD_WEATHER")
	v.AutomaticEnv()

	if x := "api-key"; v.IsSet(x) {
		o.apiKey = v.GetString(x)
	}

	if x := "units"; v.IsSet(x) {
		o.units = v.GetString(x)
	}

	if x := "location-id"; v.IsSet(x) {
		o.locationID = v.GetString(x)
	}

	if x := "city"; v.IsSet(x) {
		o.cityName = v.GetString(x)
	}

	if x := "zip-code"; v.IsSet(x) {
		o.zipcode = v.GetString(x)
	}

	return nil
}
