package go_casperjs

import (
	"bufio"
	"io/ioutil"
	"os"
	"os/exec"
	"text/template"
)

const templateName = "CasperJS"

type Casper struct {
	script         *os.File
	template       *template.Template
	bufferedWriter *bufio.Writer
	output         string
	deleteFile     bool
}

type CasperData interface {
	Execute()
}

func (c *Casper) Create() {
	var err error
	c.script, err = ioutil.TempFile(os.TempDir(), "go_casperjs_")

	if err != nil {
		panic(err)
	}

	c.template = template.New(templateName)
	if err != nil {
		panic(err)
	}
}

func (c *Casper) Open(filePath string, data CasperData) {
	template, err := template.ParseFiles(filePath)
	if err != nil {
		panic(err)
	}
	data.Execute()
	template.Execute(c.script, data)
}

func (c *Casper) Close() {
	c.script.Close()
	os.Remove(c.script.Name())
}

func (c *Casper) Run() {
	out, err := exec.Command("casperjs", c.script.Name()).Output()

	if err != nil {
		panic(err)
	}

	c.output = string(out)
}
