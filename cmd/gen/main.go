package main

import (
	"github.com/iancoleman/strcase"
	"log"
	"os"
	"text/template"

	flag "github.com/spf13/pflag"
)

type data struct {
	Name, Method, Namespace, FuncName string
}

func main() {
	var d data
	flag.StringVar(&d.Method, "method", "", "The method e.g. cancel_all_by_currency.")
	flag.StringVar(&d.Namespace, "namespace", "", "The namespace e.g. private")
	flag.Parse()
	d.Name = strcase.ToCamel(d.Method)
	d.FuncName = strcase.ToCamel(d.Namespace+d.Name)
	t := template.Must(template.New("public").Parse(methodTemplate))
	err := t.Execute(os.Stdout, d)
	if err != nil {
		log.Fatal(err)
	}
}

var methodTemplate = `

// {{.FuncName}} makes a request to {{.Namespace}}/{{.Method}}
func (e *Exchange) {{.FuncName}}(params *{{.Namespace}}.{{.Name}}Request) (*{{.Namespace}}.{{.Name}}Response, error) {
	res, err := e.rpcRequest("{{.Namespace}}/{{.Method}}", params)
	if err != nil {
		return nil, err
	}
	var ret {{.Namespace}}.{{.Name}}Response
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}`
