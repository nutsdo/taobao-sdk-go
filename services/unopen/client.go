package unopen

import "net/http"

func NewClient() *http.Client {
	client := &http.Client{
		Transport: nil,
	}

	return client
}

//func newTransport() *http.Transport {
//	return http.DefaultTransport
//}
