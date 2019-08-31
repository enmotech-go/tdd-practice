package hello

const prefixEnglishHello = "Hello, "
const prefixSpanishHello = "Hola, "
const prefixFrenchHello = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		return prefixEnglishHello + "world"
	}

	if language == "Spanish" {
		return prefixSpanishHello + name
	}

	if language == "Bonjour" {
		return prefixFrenchHello + name
	}

	return prefixEnglishHello + name
}
