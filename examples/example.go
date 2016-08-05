package main

import (
	"fmt"
	"github.com/maorsom/go-casperjs"
)

type Url struct {
	Url string
}

func (u Url) Execute() {}

func main() {

	casper := go_casperjs.Casper{}

	tpl := go_casperjs.CasperTemplate{
		TemplateFile: "/Users/maorsom/goProjects/src/github.com/maorsom/go-casperjs/examples/casper_example.js",
		Data:         Url{Url: "http://casperjs.org/"},
	}

	err := casper.Create()

	if err != nil {
		panic(err)
	}

	err = casper.LoadTemplate(tpl)

	if err != nil {
		panic(err)
	}

	err = casper.ParseString(`casper.run();`, nil)

	if err != nil {
		panic(err)
	}

	defer casper.Close()

	err = casper.Run()

	if err != nil {
		panic(err)
	}

	fmt.Println(casper.Output)
}
