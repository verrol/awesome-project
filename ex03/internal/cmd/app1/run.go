package app1

import (
	"github.com/verrol/awesome-project/ex01/featx"
	"github.com/verrol/awesome-project/ex02/featy"
)

func Run() {
	fx, _ := featx.New()

	fx.Foo()
	fx.Goo()
	fy, _ := featy.New()

	fy.Foo()
	fy.Goo()
}
