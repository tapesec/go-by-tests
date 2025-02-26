package main

import (
	"fmt"
	"net/http"
	"time"
)

type DefaultSleeper struct{}

func (s DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func main() {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Printf("Error fetching URL: %v\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)
}
