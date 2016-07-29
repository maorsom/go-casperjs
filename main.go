package main

import (
	"os"
	"io/ioutil"
	"bufio"
	"os/exec"
	"fmt"
)

type Casper struct{
	script *os.File
	bufferedWriter *bufio.Writer
	output string
	deleteFile bool
}

func (c *Casper) CreateNew(){
	var err error
	c.script, err= ioutil.TempFile(os.TempDir(), "go_casperjs_")

	if err != nil {
		panic(err)
	}

	c.deleteFile = true

	c.bufferedWriter = bufio.NewWriter(c.script)
}

func (c *Casper) Open(filePath string){
	var err error
	c.script, err = os.Open(filePath)

	if err != nil {
		panic(err)
	}

	c.deleteFile = false
}

func (c *Casper) Close(){
	c.script.Close()
	if c.deleteFile {
		os.Remove(c.script.Name())
	}
}

func (c *Casper) Write(function string){
	_, err := c.bufferedWriter.WriteString(function)

	if(err != nil){
		panic(err)
	}
}

func (c *Casper) Save() {
	c.bufferedWriter.Flush()
}

func (c *Casper) Run(){
	out, err := exec.Command("casperjs",c.script.Name()).Output();

	if(err != nil){
		panic(err)
	}

	c.output = string(out)
}

func (c *Casper) SaveAndRun(){
	c.Save()
	c.Run()
}

func main(){

	casper := Casper{}

	casper.Open(`C:\Users\Somech\Desktop\casperjs.js`)

	defer casper.Close()

	casper.Run()

	fmt.Println(casper.output)
}