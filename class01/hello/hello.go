package hello

const prefixEnglishHello = "Hello, "

func Hello(name string) string {
	if name == "" {
		return prefixEnglishHello + "world"
	}

	return prefixEnglishHello + name
}
