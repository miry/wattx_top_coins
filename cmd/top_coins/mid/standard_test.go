package mid_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/miry/wattx_top_coins/cmd/top_coins/app"
	"github.com/miry/wattx_top_coins/cmd/top_coins/mid"
	"github.com/stretchr/testify/assert"
)

// GetTestPanicHandler returns a http.HandlerFunc for testing http middleware
func GetTestPanicHandler(w http.ResponseWriter, r *http.Request) {
	panic("it should be handle")
}

func TestPanic(t *testing.T) {
	assert := assert.New(t)

	app, err := app.NewApp()
	assert.NotNil(t, err)

	ts := httptest.NewServer(http.HandlerFunc(mid.PanicMiddleware(app, GetTestPanicHandler)))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	assert.NotNil(t, err)

	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	assert.NotNil(t, err)

	assert.Equal(500, res.StatusCode)
	assert.Equal("it should be handle\n", string(result))
}
