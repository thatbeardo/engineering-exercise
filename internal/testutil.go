package internal

import (
	"net/http"
	"net/http/httptest"
	"strings"
)

func PerformRequest(r http.Handler, method, path, body string) (*http.Response, func() error) {
	req, _ := http.NewRequest(method, path, strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, req)
	response := recorder.Result()
	return response, response.Body.Close
}
