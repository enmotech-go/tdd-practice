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
	if language == spanish {
		return prefixSpanish + name
	}
	if language == french {
		return prefixFrench + name
	}
	return helloPrefix + name
}
