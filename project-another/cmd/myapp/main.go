package main

import (
	"log"

	"github.com/verrol/awesome-project/ex03/featx"
)

func main() {
	fx, err := featx.New()

	if err != nil {
		log.Fatal(err)
	}

	fx.Foo()
	fx.Goo()
}
