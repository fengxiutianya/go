package test

import (
	"log"
	"os"
	"strings"
	"testing"
	"text/template"
)

/**************************************
 *      Author : zhangke
 *      Date   : 2019-04-14 16:26
 *      email  : 398757724@qq.com
 *      Desc   :
 ***************************************/

func Test_study1(t *testing.T) {

	type Inventory struct {
		Material string
		Count    uint
	}

	sweaters := Inventory{"wool", 17}

	tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}\n")

	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, sweaters)

	if err != nil {
		panic(err)
	}
}

func Test_demo1(t1 *testing.T) {
	const letter = `
		Dear {{.Name}},
		{{if .Attended}}
		It was a pleasure to see you at the wedding.
		{{- else}}
		It is a shame you couldn't make it to the wedding.
		{{- end}}
		{{with .Gift -}}
		Thank you for the lovely {{.}}.
		{{end}}
		Best wishes,
		Josie

		`

	// Prepare some data to insert into the template.
	type Recipient struct {
		Name, Gift string
		Attended   bool
	}
	var recipients = []Recipient{
		{"Aunt Mildred", "bone china tea set", true},
		{"Uncle John", "moleskin pants", false},
		{"Cousin Rodney", "", false},
	}

	// Create a new template and parse the letter into it.
	t := template.Must(template.New("letter").Parse(letter))

	// Execute the template for each recipient.
	for _, r := range recipients {
		err := t.Execute(os.Stdout, r)
		if err != nil {
			log.Println("executing template:", err)
		}
	}
}
func Test_block(t *testing.T) {
	const (
		master  = `Names:{{block "list" .}}{{"\n"}}{{range .}}{{println "-" .}}{{end}}{{end}}`
		overlay = `{{define "list"}} {{join . ", "}}{{end}}`
	)
	var (
		funcs     = template.FuncMap{"join": strings.Join}
		guardians = []string{"Gamora", "Groot", "Nebula", "Rocket", "Star-Lord"}
	)
	masterTmpl, err := template.New("master").Funcs(funcs).Parse(master)
	if err != nil {
		log.Fatal(err)
	}

	overlayTmpl, err := template.Must(masterTmpl.Clone()).Parse(overlay)
	if err != nil {
		log.Fatal(err)
	}

	if err := masterTmpl.Execute(os.Stdout, guardians); err != nil {
		log.Fatal(err)
	}

	if err := overlayTmpl.Execute(os.Stdout, guardians); err != nil {
		log.Fatal(err)
	}
}
func Test_actions(t *testing.T) {
	test1, err := template.New("test1").Parse(" {{- 23  -}}1233 {{- /* a comment with white space trimmed from preceding and following text */ -}}");
	if err != nil {
		t.Fatal(err)
	}
	test1.Execute(os.Stdout, nil)
}

func Test_print(t *testing.T) {
	h1 := `
{{range $key,$value := .}}{{$key}}{{"."}} {{$value}}{{"\n"}}{{end}}

`
	file, err := os.OpenFile("test.md", os.O_APPEND|os.O_RDWR|os.O_CREATE, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	test1, _ := template.New("h1").Parse(h1);
	test := []string{"dd", "dddd"}

	test1.Execute(file, test)
}
