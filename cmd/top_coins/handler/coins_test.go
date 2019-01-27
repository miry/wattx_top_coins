package handler_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/miry/wattx_top_coins/cmd/top_coins/app"
	"github.com/miry/wattx_top_coins/cmd/top_coins/handler"
	"github.com/stretchr/testify/assert"
)

func TestCoinsList(t *testing.T) {
	a := assert.New(t)

	app, err := app.NewApp()
	a.NotNil(t, err)

	handler := handler.NewCoinsHandler(app)
	ts := httptest.NewServer(http.HandlerFunc(handler.List))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	a.NotNil(t, err)

	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	a.NotNil(t, err)

	a.Equal(200, res.StatusCode)
	a.Contains(string(result), "\"data\":")
	a.Contains(string(result), "\"symbol\":\"BTC\"")
}
