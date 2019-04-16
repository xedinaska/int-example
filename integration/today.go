package integration

import (
	"context"

	"net/http"

	"github.com/pkg/errors"
	"github.com/xedinaska/int-weather-sdk/api"
)

type TodayWeather struct {
	StateName  string  `json:"state_name"`
	MinTemp    float64 `json:"min_temp"`
	MaxTemp    float64 `json:"max_temp"`
	WindSpeed  float64 `json:"wind_speed"`
	Humidity   int     `json:"humidity"`
	Visibility float64 `json:"visibility"`
	Success    bool    `json:"success"`
}

func (i *Example) GetTodayWeather(ctx context.Context, req *api.TodayWeatherRequest) (*api.TodayWeatherResponse, error) {

	target := &TodayWeather{}
	call, err := i.RequestClient.Post(ctx, "/", req, target)

	if err != nil {
		i.Logger.WithField("call", call).Errorf("failed to get weather for today")
		return nil, err
	}

	if call.ResponseStatus != http.StatusOK || !target.Success {
		return nil, errors.New("response from server isn't success")
	}

	response := &api.TodayWeatherResponse{
		StateName:  target.StateName,
		MinTemp:    target.MinTemp,
		MaxTemp:    target.MaxTemp,
		Humidity:   target.Humidity,
		WindSpeed:  target.WindSpeed,
		Visibility: target.Visibility,
	}

	return response, nil
}
