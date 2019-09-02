package hello_world

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
	str = greetingPrefix(language)

	return  str+ name
}

type a struct {
	name string
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = FrenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = helloPrefix
	}
	return
}