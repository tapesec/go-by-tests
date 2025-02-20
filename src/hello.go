package main

import "fmt"

func main() {
	fmt.Println(Hello("Chris"))
}

const greetingPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix + name
}
