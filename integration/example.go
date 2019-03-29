package integration

import (
	"crypto/tls"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/xedinaska/int-example/request"
)

type Example struct {
	Logger        *logrus.Entry
	RequestClient *request.RequestClient
}

func Init(logger *logrus.Entry) *Example {
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	return &Example{
		Logger: logger,
		RequestClient: &request.RequestClient{
			HTTPClient: httpClient,
		},
	}
}
