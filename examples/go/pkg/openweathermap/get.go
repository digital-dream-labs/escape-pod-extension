package openweathermap

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	endpoint           = "api.openweathermap.org/data/2.5"
	httpTimeoutSeconds = 2
)

// response is a truncated list of what we actually need.  For full details, see:
// https://openweathermap.org/current#current_JSON
type response struct {
	Weather []struct {
		ID int `json:"id"`
	} `json:"weather"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Name string `json:"name"`
}

// Get returns the map of response data from openweathermap
func (o *OpenWeatherMap) Get(ctx context.Context, arg string, m map[string]string) (string, map[string]string, error) {

	client := &http.Client{
		Timeout: time.Second * httpTimeoutSeconds,
	}

	resp, err := client.Get(
		fmt.Sprintf(
			"http://%v/weather?%v&units=%v&appid=%v",
			endpoint,
			o.location,
			o.units,
			o.apiKey,
		),
	)
	if err != nil {
		return "", nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	var r response
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return "", nil, err
	}

	var condition string
	cond, ok := conditionsMap[r.Weather[0].ID]
	if !ok {
		condition = "Not Available (N/A)"
	} else {
		condition = cond
	}

	data := map[string]string{
		"speakable_location_string": r.Name,
		"is_forecast":               "false", // forecast not supported yet
		"condition":                 condition,
		"temperature":               fmt.Sprintf("%.0f", r.Main.Temp),
		"temperature_unit":          o.getUnit(),
	}

	return arg, data, nil

}

func (o *OpenWeatherMap) getUnit() string {
	if o.units == "imperial" {
		return "F"
	}
	return "C"
}
