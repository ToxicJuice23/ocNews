package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func curl(url string, timeout time.Duration) *http.Response {
	client := &http.Client{Timeout: timeout}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Errorf("Error: %v\n", err)
		os.Exit(1)
	}
	r, err := client.Do(req)
	if err != nil {
		fmt.Errorf("Error: %v\n", err)
		os.Exit(1)
	}
	return r
}
