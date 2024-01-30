package main

import "fmt"

const (
    spanish = "Spanish"
    french = "French"
    portuguese = "Portuguese"
    japanese = "Japanese"

    englishHelloPrefix = "Hello, "
    spanishHelloPrefix = "Hola, "
    frenchHelloPrefix = "Bonjour, "
    portugueseHelloPrefix = "Olá, "
    japaneseHelloPrefix = "Kon'nichiwa, "
)

func Hello(name string, lang string) string {
    if name ==  "" {
        name = "World"
    }
    return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) (prefix string) {
    switch lang {
        case spanish: prefix = spanishHelloPrefix
        case french: prefix = frenchHelloPrefix
        case portuguese: prefix = portugueseHelloPrefix
        case japanese: prefix = japaneseHelloPrefix
        default: prefix = englishHelloPrefix
    }
    return
}

func main() {
    fmt.Println(Hello("World!", ""));
}
