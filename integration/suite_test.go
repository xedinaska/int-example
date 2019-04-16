package integration_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"github.com/xedinaska/int-example/config"
	"github.com/xedinaska/int-example/integration"
)

type IntegrationTestSuite struct {
	suite.Suite
	ctx          context.Context
	app          *integration.Example
	ts           *httptest.Server
	mockResponse []byte
}

func (i *IntegrationTestSuite) SetupTest() {
	logger := logrus.WithFields(logrus.Fields{
		"logger": "test",
	})

	i.app = integration.Init(logger)
	i.ctx = context.Background()
}

func (i *IntegrationTestSuite) TeardownTest() {
	i.ts.Close()
}

func (i *IntegrationTestSuite) configureAndStartTestServer(routeMap map[string]string) error {
	mux := http.NewServeMux()
	for route, filepath := range routeMap {
		data, err := ioutil.ReadFile(filepath)
		if err != nil {
			return errors.Wrapf(err, "could not open testdata file '%s'", filepath)
		}

		mux.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, string(data))
		})
	}

	i.ts = httptest.NewServer(mux)
	i.ctx = context.WithValue(i.ctx, config.BaseURL, i.ts.URL)

	return nil
}
