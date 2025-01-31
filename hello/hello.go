package hello

import "fmt"

const (
	english = "English"
	spanish = "Spanish"

	french             = "French"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) string {
	prefix := englishHelloPrefix
	switch lang {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	}
	return prefix
}

func main() {
	fmt.Println(Hello("World", "English"))
}
