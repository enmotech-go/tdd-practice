package class1

const helloPrefix = "Hello, "
const PrefixSpanish = "Hola, "

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}
	if language == "Spanish" {
		return PrefixSpanish + name
	}
	return helloPrefix + name
}
