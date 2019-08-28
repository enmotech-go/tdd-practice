package hello

const (
	langEnglish = "English"
	langChinese = "Chinese"
)

const (
	helloInEnglish = "Hello, "
	helloInChinese = "你好，"
)

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}

	return getHelloWord(lang) + name
}

func getHelloWord(lang string) string {
	var word string

	switch lang {
	case langEnglish:
		word = helloInEnglish
	case langChinese:
		word = helloInChinese
	default:
		word = helloInEnglish
	}
	return word
}
