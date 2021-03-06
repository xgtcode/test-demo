package test_demo

import "fmt"

type Option struct {
	Name string
}

type Options func(*Option)

var defaultOption = &Option{}

func NewOptions(options ...Options) *Option{
	for _, opts := range options{
		opts(defaultOption)
	}
	return defaultOption
}

func WithNameOption(name string) Options{
	return func(opt *Option){
		opt.Name = name
	}
}
func PrintDefaultOption() {
	fmt.Println(defaultOption)
}
