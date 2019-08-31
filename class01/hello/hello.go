package hello

const englishPrefix = "Hello, "
const spanishPrefix = "Hola, "
const frenchPrefix = "Bonjour, "
const spanish = "Spanish"
const french = "French"

func generatePrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchPrefix
	case spanish:
		prefix = spanishPrefix
	default:
		prefix = englishPrefix
	}
	return
}
func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	prefix := generatePrefix(language)

	return prefix + name
}
