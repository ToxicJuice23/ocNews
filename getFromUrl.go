package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func getFromUrl(url string, timeout time.Duration) *http.Response {
	client := &http.Client{Timeout: timeout}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		_ = fmt.Errorf("%v", err.Error())
		os.Exit(1)
	}
	r, err := client.Do(req)
	if err != nil {
		_ = fmt.Errorf("%v", err.Error())
		os.Exit(1)
	}
	return r
}
