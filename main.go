package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

type requestResult struct {
	url    string
	status string
}

var errRequestFailed = errors.New("Request failed")

func main() {
	// 초기화하지 않고 아래와 같이 작성하면 nil, 값을 넣을 수 없음
	// var results map[string]string
	// panic: assignment to entry in nil map

	// 또는 make 함수로 생성
	results := map[string]string{}
	c := make(chan requestResult)
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
		go hitURL(url, c)
	}
	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}
	for url, status := range results {
		fmt.Println(url, status)
	}

	// c := make(chan string)
	// people := [5]string{"nico", "flynn", "a", "b", "c"}
	// for _, person := range people {
	// 	go isSexy(person, c)

	// }
	// for i := 0; i < len(people); i++ {

	// 	fmt.Println(<-c)
	// }
}

func hitURL(url string, c chan<- requestResult) {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAIL"
	}
	c <- requestResult{url: url, status: status}
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 3)
	c <- person + " is sexy"
}
