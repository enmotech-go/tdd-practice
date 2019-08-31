package hello

const prefixEnglishHello = "Hello, "
const prefixSpanishHello = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		return prefixEnglishHello + "world"
	}

	if language == "Spanish" {
		return prefixSpanishHello + name
	}

	return prefixEnglishHello + name
}
