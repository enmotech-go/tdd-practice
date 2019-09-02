package main

import "fmt"

const (
	helloPrefix = "Hello, "
	spanish = "Spanish"
	french = "French"
	spanishHelloPrefix = "Hola, "
	FrenchHelloPrefix = "Bonjour, "
)
func Hello(name string, language string) (str string) {
	if name == "" {
		name = "World"
	}
	str,ll := greetingPrefix(language)
	fmt.Print(ll)
	return  str+ name
}

type a struct {
	name string
}

func greetingPrefix(language string) (prefix string,mm string) {
	switch language {
	case french:
		prefix = FrenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = helloPrefix
	}
	return "xx","mm"
}