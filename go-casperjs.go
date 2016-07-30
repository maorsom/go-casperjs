package go_casperjs

import (
	"io/ioutil"
	"os"
	"os/exec"
	"text/template"
)

const (
	templateName = "CasperJS"
	commandName = "casperjs"
)


type Casper struct {
	script         *os.File
	template       *template.Template
	Output         string
}

type CasperData interface {
	Execute()
}

type CasperTemplate struct {
	Name string
	Dir string
	Data CasperData
}

func (tpl *CasperTemplate) GetPath() string {
	return tpl.Dir + "\\" + tpl.Name
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

func (c *Casper) LoadTemplate(cTemplate CasperTemplate) {

	template, err := template.ParseFiles(cTemplate.GetPath())
	if err != nil {
		panic(err)
	}
	if cTemplate.Data != nil {
		cTemplate.Data.Execute()
	}

	template.Execute(c.script, cTemplate.Data)
}

func (c *Casper) ParseString(content string,data CasperData){
	var err error
	c.template, err = c.template.Parse(content)
	if err != nil {
		panic(err)
	}

	err = c.template.Execute(c.script, data)
	if err != nil {
		panic(err)
	}
}

func (c *Casper) Close() {
	c.script.Close()
	os.Remove(c.script.Name())
}

func (c *Casper) Run() {
	out, err := exec.Command(commandName, c.script.Name()).Output()

	if err != nil {
		panic(err)
	}

	c.Output = string(out)
}
