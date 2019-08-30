package helloworld

//import "fmt"
const (
	helloPrefix  = "Hello "
	spanishHello = "Hola, "
	frenchHello  = "Bonjour, "
)

func hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "French":
		prefix = frenchHello
	case "Spanish":
		prefix = spanishHello
	default:
		prefix = helloPrefix
	}
	return
}

//func hello(s string, language string) string {
//	if s == "" {
//		s = "World"
//	}
//	prefix := helloPrefix
//	switch language {
//	case "Spanish":
//		prefix = spanishHello
//	case "French":
//		prefix = frenchHello
//	}
//	return prefix + s
//
//}
