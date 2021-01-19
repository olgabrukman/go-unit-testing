package example5_httptest

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	url   = "/health-check"
	alive = `{"alive": true}`
)

func TestHealthCheckHandlerGETFull(t *testing.T) {
	handler := http.HandlerFunc(HealthCheckHandler)
	req := httptest.NewRequest("GET", url, nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	//must close body!
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "", resp.Header.Get("Content-Type"))
	assert.Equal(t, alive, string(body))
}

func TestHealthCheckHandlerGETShort(t *testing.T) {
	handler := http.HandlerFunc(HealthCheckHandler)

	assert.HTTPSuccess(t, handler, "GET", url, nil)
	assert.HTTPBodyContains(t, handler, "GET", url, nil, alive)
}
