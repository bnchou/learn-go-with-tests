package helloworld

// Domain code
func Hello(name, language string) string {
	const englishHelloPrefix = "Hello, "
	const spanishHelloPrefix = "Hola, "

	if name == "" {
		name = "World"
	}

	if language == "Spanish" {
		return spanishHelloPrefix + name
	}

	return englishHelloPrefix + name
}
