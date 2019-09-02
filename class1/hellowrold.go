package class1

const helloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "world"
	}
	return helloPrefix + name
}
