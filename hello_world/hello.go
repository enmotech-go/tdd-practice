package main

import (
	"fmt"
)

/**
class_01

demand: 1. 问候某人; 2. 制定问候的语言
*/

const (
	englishHelloPrefix = "Hello, "
	spanish            = "Spanish"
	spanishHelloPrefix = "Hola, "
	french             = "French"
	frenchHelloPrefix  = "French, "
)

func main() {
	fmt.Println(Hello("world", ""))
}

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}
	// if language == spanish {
	// 	return spanishHelloPrefix+ name
	// }
	// if language == french {
	//     return frenchHelloPrefix + name
	// }
	// name = englishHelloPrefix + name

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case frenchHelloPrefix:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
