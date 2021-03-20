package main

import (
	"flag"
	"fmt"
)

// common CLI lib: https://github.com/urfave/cli

// флаги в Go имеют стиль Plan 9 - без -- и склейки однобуквенных флагов :/
// 2 варианта флагов - однозначные (значение парсится в указатель)
// и с псевдонимами (значение парсится в переменную)
var name = flag.String("name", "World", "A name to say hello to.")
var spanish bool

// регистрацию и парсинг флагов удобно выносить в init.
func init() {
	flag.BoolVar(&spanish, "spanish", false, "Use Spanish language.")
	flag.BoolVar(&spanish, "s", false, "Use Spanish language.")
	flag.Parse()
}

func main() {
	if spanish == true {
		fmt.Printf("Hola %s!\n", *name)
	} else {
		fmt.Printf("Hello %s!\n", *name)
	}
}
