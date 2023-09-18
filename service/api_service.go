package service

import (
	"io"
	"net/http"
)

type Service struct {
	body io.ReadCloser
}

func createAPI(service *Service) *http.Response {
	respone := &Service{}
	var url = "localhost:8080/login"
	http.Post(url, "application/json", http.NoBody)
	return (&http.Response{
		StatusCode: 200,
		Body:       respone.body,
	})
}
