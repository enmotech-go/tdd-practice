package class1

const helloPrefix = "Hello, "
const prefixSpanish = "Hola, "
const spanish = "Spanish"
const french = "French"
const prefixFrench = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) string {
	switch language {
	case spanish:
		return prefixSpanish
	case french:
		return prefixFrench
	default:
		return helloPrefix
	}
}
