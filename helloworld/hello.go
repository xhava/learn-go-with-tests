package main

import "fmt"

const spanish = "spanish"
const french = "french"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {

	prefix := englishHelloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
		if name == "" {
			name = "Mundo"
		}
	case french:
		prefix = frenchHelloPrefix
		if name == "" {
			name = "Monde"
		}
	}

	if name == "" {
		name = "World"
	}

	return prefix + name
}
func main() {
	fmt.Println(Hello("Chrystopher", "french"))
}
