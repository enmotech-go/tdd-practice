package class1

const helloPrefix = "Hello, "
const prefixSpanish = "Hola, "
const spanish = "Spanish"

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}
	if language == spanish {
		return prefixSpanish + name
	}
	return helloPrefix + name
}
