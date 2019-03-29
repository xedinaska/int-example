package integration

import (
	"context"

	"github.com/pkg/errors"
	"github.com/xedinaska/int-weather-sdk/api"
)

func (i *Example) GetTodayWeather(ctx context.Context, req *api.TodayWeatherRequest) (*api.TodayWeatherResponse, error) {
	return nil, errors.New("not implemented")
}
