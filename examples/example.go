package main

import (
	"fmt"
	"github.com/maorsom/go-casperjs"
	"os"
)

type Url struct {
	Url string
}

func (u Url) Execute() {}

func main(){

	casper := go_casperjs.Casper{}

	dir, _ := os.Getwd()

	tpl := go_casperjs.CasperTemplate{
		Dir: dir + `\examples`,
		Name : "casper_example.js",
		Data : Url{Url:"http://casperjs.org/"},
	}

	casper.Create()
	casper.LoadTemplate(tpl)
	casper.ParseString(`casper.run();`,nil)

	defer casper.Close()

	casper.Run()

	fmt.Println(casper.Output)
}
