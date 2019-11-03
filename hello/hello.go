package main

import "fmt"

const spanishLanguage = "Spanish"
const frenchLanguage = "French"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchhHelloPrefix = "Bonjur, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case frenchLanguage:
		prefix = frenchhHelloPrefix
	case spanishLanguage:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
