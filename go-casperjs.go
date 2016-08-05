package go_casperjs

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

const (
	templateName = "CasperJS"
	commandName  = "casperjs"
)

type Casper struct {
	script   *os.File
	template *template.Template
	Output   string
}

type CasperData interface {
	Execute()
}

type CasperTemplate struct {
	TemplateFile string
	Data         CasperData
}

func (tpl *CasperTemplate) GetPath() string {
	return filepath.FromSlash(tpl.TemplateFile)
}

func (c *Casper) Create() error {

	//create new temp file inside temp folder. casperjs will ran this file.
	var err error
	c.script, err = ioutil.TempFile(os.TempDir(), "go_casperjs_")

	if err != nil {
		return errors.New("GoCasperjs:: Can't create temp file" + err.Error())
	}

	//create new template, throw erorr if can't
	c.template = template.New(templateName)

	return nil //return nil - no error.
}

func (c *Casper) LoadTemplate(cTemplate CasperTemplate) error {

	//load template files.
	template, err := template.ParseFiles(cTemplate.GetPath())
	if err != nil {
		return errors.New("GoCasperjs:: can't parse template files: " + err.Error())
	}

	//if data not supplied don't execute data.
	if cTemplate.Data != nil {
		cTemplate.Data.Execute()
	}

	//execute template and add into script.
	err = template.Execute(c.script, cTemplate.Data)
	if err != nil {
		return errors.New("GoCasperjs:: cant execute template with data: " + err.Error())
	}

	return nil

}

func (c *Casper) ParseString(content string, data CasperData) error {
	var err error
	c.template, err = c.template.Parse(content)
	if err != nil {
		return errors.New("GoCasper:: can't parse string with data: " + err.Error())
	}

	err = c.template.Execute(c.script, data)
	if err != nil {
		return errors.New("GoCasperjs:: can't execute string with data: " + err.Error())
	}

	return nil
}

func (c *Casper) Close() {
	c.script.Close()
	os.Remove(c.script.Name())
}

func (c *Casper) Run() error {
	out, err := exec.Command(commandName, c.script.Name()).Output()

	if err != nil {
		return errors.New("GoCasperjs:: failed to run casperjs: " + err.Error())
	}

	c.Output = string(out)

	return nil
}
