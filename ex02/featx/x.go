package featx

import "fmt"

type FeatureX struct{}

func New() (*FeatureX, error) {
	return &FeatureX{}, nil
}

func (f *FeatureX) Foo() {
	fmt.Println("FeatureX.Foo()")
}

func (f *FeatureX) Goo() {
	fmt.Println("FeatureX.Goo()")
}
