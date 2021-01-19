package openweathermap

import (
	"fmt"
)

// OpenWeatherMap is a configuration struct
type OpenWeatherMap struct {
	apiKey   string
	location string
	units    string
}

// New returns a populated openweathermap struct
func New(opts ...Option) (*OpenWeatherMap, error) {
	cfg := options{}

	for _, opt := range opts {
		opt(&cfg)
	}

	if cfg.apiKey == "" {
		return nil, fmt.Errorf("api key must be set")
	}

	if cfg.units == "" || cfg.units != "imperial" && cfg.units != "metric" {
		return nil, fmt.Errorf("a valid unit must be set")
	}

	ow := OpenWeatherMap{
		apiKey: cfg.apiKey,
		units:  cfg.units,
	}

	switch {
	case cfg.locationID != "" && cfg.cityName == "" && cfg.zipcode == "":
		ow.location = fmt.Sprintf("id=%v", cfg.locationID)
	case cfg.locationID == "" && cfg.cityName != "" && cfg.zipcode == "":
		ow.location = fmt.Sprintf("q=%v", cfg.cityName)
	case cfg.locationID == "" && cfg.cityName == "" && cfg.zipcode != "":
		ow.location = fmt.Sprintf("zip=%v", cfg.zipcode)
	default:
		return nil, fmt.Errorf("invalid location configuration.  please choose either location id, city name, or zip code")
	}

	return &ow, nil
}
