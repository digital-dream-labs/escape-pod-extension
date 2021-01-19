# openweathermap

An escape pod interface to openweathermap.orgs service.  

## Usage

This is meant to interface with other packages in this repository.  Feel free to use it elsewhere if you find it useful.

## Environment variables

|Variable| Description |
|--|--|
| ESCAPE_POD_WEATHER_API_KEY | Your API key from https://openweathermap.org/api |
| ESCAPE_POD_WEATHER_UNITS | Temperature units (metric or imperial) |
| ESCAPE_POD_WEATHER_LOCATION_ID | Your location ID |
| ESCAPE_POD_WEATHER_CITY | Your city |
| ESCAPE_POD_WEATHER_ZIP_CODE | Your zip code |

  Note:  Only set one of the location_id, city, or zip code.  Any one of them will work, but multiples will not.

## Initalization options

Please see options.go for a full list of options