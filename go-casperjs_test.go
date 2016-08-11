package go_casperjs

import "testing"

func TestCreate(t *testing.T) {
	casper := Casper{}

	err := casper.Create()
	if err != nil {
		t.Fatalf("got an error %s", err.Error())
	}

	if casper.script == nil {
		t.Fatal("Unable to create script file")
	}

	if casper.template == nil {
		t.Fatal("Unable to create template file")
	}
}

func TestGetPath(t *testing.T) {

	filepath := "/Users/user/Desktop/example.js"

	template := CasperTemplate{
		TemplateFile: filepath,
		Data:         nil,
	}

	newpath := template.GetPath()

	if newpath != filepath {
		t.Fatalf("expected %s but got %s", newpath, filepath)
	}

}
