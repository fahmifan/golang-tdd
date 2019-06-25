package greeting

const (
	spanish = "spanish"
	french  = "french"

	spanishPrefix = "Hola, "
	englishPrefix = "Hello, "
	frenchPrefix  = "Bonjour, "
)

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	default:
		prefix = englishPrefix
	}

	return prefix
}

// Hello method greeting people
func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(lang) + name
}
