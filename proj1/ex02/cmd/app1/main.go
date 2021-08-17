package main

import (
	"github.com/verrol/awesome-project/proj1/ex02/featx"
	"github.com/verrol/awesome-project/proj1/ex02/featy"
)

func main() {
	fx, _ := featx.New()

	fx.Foo()
	fx.Goo()
	fy, _ := featy.New()

	fy.Foo()
	fy.Goo()
}
