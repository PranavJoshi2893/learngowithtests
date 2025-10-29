package hello

const (
	french  = "French"
	spanish = "Spanish"

	greetingEnglishPrefix = "Hello, "
	greetingFrenchPrefix  = "Bonjour, "
	greetingSpanishPrefix = "Hola, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefixSelector(language) + name
}

func greetingPrefixSelector(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = greetingSpanishPrefix
	case french:
		prefix = greetingFrenchPrefix
	default:
		prefix = greetingEnglishPrefix
	}
	return
}
