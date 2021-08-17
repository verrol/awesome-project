package featy

import "fmt"

type FeatureY struct{}

func New() (*FeatureY, error) {
	return &FeatureY{}, nil
}

func (f *FeatureY) Foo() {
	fmt.Println("FeatureY.Foo()")
}

func (f *FeatureY) Goo() {
	fmt.Println("FeatureY.Goo()")
}
