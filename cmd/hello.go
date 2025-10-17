package main

import "fmt"

const en = "hi, "
const tr = "merhaba, "
const es = "gracias, "

func Hello(name, language string) string {
	if name == "" {
		name = "SUPRA"
	}
	if language == "" {
		language = "en"
	}

	switch language {
	case "en":
		return en + name
	case "tr":
		return tr + name
	case "es":
		return es + name
	default:
		return en + name
	}
}

func main() {
	fmt.Println(Hello("world", "Turkish"))
}
