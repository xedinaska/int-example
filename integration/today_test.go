package integration_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/xedinaska/int-weather-sdk/api"
)

type TodayWeatherTestSuite struct {
	IntegrationTestSuite
}

func (s *TodayWeatherTestSuite) TestGetTodayWeather_Success() {
	routeMap := map[string]string{
		"/": "testdata/today_weather_success.json",
	}
	err := s.configureAndStartTestServer(routeMap)
	s.Require().NoError(err)

	response, err := s.app.GetTodayWeather(s.ctx, &api.TodayWeatherRequest{
		Latitude:  42.21,
		Longitude: 25.11,
	})

	s.Require().NoError(err)
	s.Require().Equal("sunny", response.StateName)
	s.Require().Equal(10.00, response.MinTemp)
	s.Require().Equal(12.30, response.MaxTemp)
	s.Require().Equal(30, response.Humidity)
	s.Require().Equal(22.1, response.WindSpeed)
	s.Require().Equal(90.11, response.Visibility)
}

func (s *TodayWeatherTestSuite) TestGetTodayWeather_Failure() {
	routeMap := map[string]string{
		"/": "testdata/today_weather_failed.json",
	}
	err := s.configureAndStartTestServer(routeMap)
	s.Require().NoError(err)

	_, err = s.app.GetTodayWeather(s.ctx, &api.TodayWeatherRequest{
		Latitude:  42.21,
		Longitude: 25.11,
	})

	s.Require().NotNil(err)
	s.Require().Equal("response from server isn't success", err.Error())
}

func TestTodayWeather(t *testing.T) {
	suite.Run(t, new(TodayWeatherTestSuite))
}
