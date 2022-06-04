package hello

import "fmt"

const Spanish = "Spanish"
const French = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

// Hello returns a personalised greeting in a given language.
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return GreetingPrefix(language) + name
}

func GreetingPrefix(language string) (prefix string) {
	switch language {
	case French:
		prefix = frenchHelloPrefix
	case Spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func Mainly() {
	fmt.Println(Hello("world", ""))
}
