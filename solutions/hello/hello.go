package main

import (
	"fmt"
)

const french  = "french"
const spanish  = "spanish"
const chinese  =  "chinese"
const englishPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const chineseHelloPrefix  = "你好, "


func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return Greeting(language) + name

}

func Greeting(language string) (prefix string)  {

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case chinese:
		prefix = chineseHelloPrefix
	default:
		prefix = englishPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
