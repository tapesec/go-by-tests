package main

import "fmt"

func main() {
	fmt.Println(Hello("Chris", ""))
}

const (
	englishGreetingPrefix = "Hello, "
	spanishGreetingPrefix = "Hola, "
	frenchGreetingPrefix  = "Bonjour, "
	french                = "French"
	spanish               = "Spanish"
)

func Hello(name, country string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(country) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchGreetingPrefix
	case spanish:
		prefix = spanishGreetingPrefix
	default:
		prefix = englishGreetingPrefix
	}
	return
}
