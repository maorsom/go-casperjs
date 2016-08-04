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
		Dir:  "/Users/maorsom/goProjects/src/github.com/maorsom/go-casperjs/example",
		Name: "casper_example.js",
		Data: Url{Url: "http://casperjs.org/"},
	}

	casper.Create()
	casper.LoadTemplate(tpl)
	casper.ParseString(`casper.run();`, nil)

	defer casper.Close()

	casper.Run()

	fmt.Println(casper.Output)
}
