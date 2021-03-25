package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const giryama = "Giryama"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const giryamaHelloPrefix = "Jambo, "

func Hello(name string, language string) string {
    if name == "" {
        name = "World"
    }
    return greetingFrefix(language) + name

}

func greetingFrefix(language string) (prefix string) {
    switch language {
        case french:
            prefix = frenchHelloPrefix
        case spanish:
            prefix = spanishHelloPrefix
        case giryama:
            prefix = giryamaHelloPrefix
        default:
            prefix = englishHelloPrefix
    }
    return
}

func main() {
    fmt.Println(Hello("world", "English"))
}
