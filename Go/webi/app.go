package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

var templateP *template.Template

var funcMap template.FuncMap

func increment(v int, inc int) int {
	return v + inc
}

func initTemplates(templatePathPattern string) {

	funcMap = template.FuncMap{
		"UC":  strings.ToUpper,
		"INC": increment,
	}
	var err error
	templateP = template.New("")
	templateP.Funcs(funcMap)
	templateP, err = templateP.ParseGlob(templatePathPattern)
	if err != nil {
		log.Fatal("Template initiation error -> ", err)
	}
}

func main() {
	var err error
	data := []struct {
		FirstName string
		LastName  string
		Age       int
	}{{"tej", "kalagara", 23}, {"gokul", "kalagara", 20}}
	initTemplates("templates/*.gohtml")
	fmt.Println("HOME PAGE -------------------------------")
	err = templateP.ExecuteTemplate(os.Stdout, "home.gohtml", data)
	if err != nil {
		log.Fatal(err)
	}

}
