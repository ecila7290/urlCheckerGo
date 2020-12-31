package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request failed")

func main() {
	// 초기화하지 않고 아래와 같이 작성하면 nil, 값을 넣을 수 없음
	// var results map[string]string
	// panic: assignment to entry in nil map

	// 또는 make 함수로 생성
	results := map[string]string{}
	urls := []string{
		"https://www.airbnb.com",
		"https://www.google.com",
		"https://www.amazon.com",
		"https://www.reddit.com",
		"https://www.soundcloud.com",
		"https://www.facebook.com",
		"https://www.instagram.com",
	}

	for _, url := range urls {
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}
	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string) error {
	fmt.Println("Checking ", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err, resp.StatusCode)
		return errRequestFailed
	}
	return nil
}