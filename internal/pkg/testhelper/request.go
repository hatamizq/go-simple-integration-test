package testhelper

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/go-chi/chi"
)

func BuildRequest(method, url string, body any) *http.Request {
	jsonBody, _ := json.Marshal(body)
	reqBody := bytes.NewReader(jsonBody)
	req, _ := http.NewRequest(method, url, reqBody)
	return req
}

func SendRawHTTPTestRequest(router chi.Router, request *http.Request) *http.Response {
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)
	return w.Result()
}
