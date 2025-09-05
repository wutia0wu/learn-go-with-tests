package main

import (
	"fmt"
)

// const
const spanish = "Spanish"
const chinese = "Chinese"
const englishHelloPrefix = "Hello, "
const chineseHelloPrefix = "你好, "
const spanishHelloPrefix = "Hola, "

func Hello(name string, language string) string {

	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case chinese:
		prefix = chineseHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("Gopher", ""))
}
