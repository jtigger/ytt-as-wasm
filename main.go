package main

import (
	"github.com/k14s/ytt/pkg/files"
	"syscall/js"
	cmdtpl "github.com/k14s/ytt/pkg/cmd/template"
	"fmt"
    cmdcore "github.com/k14s/ytt/pkg/cmd/core"
)

func add(this js.Value, args []js.Value) interface{} {
	sum := args[0].Int() + args[1].Int()
	js.Global().Set("output", js.ValueOf(sum))
	fmt.Printf("%d + %d = %d", args[0].Int(), args[1].Int(), sum)
	return sum
}

func doTemplate(this js.Value, args []js.Value) interface{} {
	yamlTplData := []byte(`
#@ load("@ytt:data", "data")
data_int: #@ data.values.int
data_str: #@ data.values.str`)

	yamlData := []byte(`
#@data/values
---
int: 123
str: str`)

	filesToProcess := files.NewSortedFiles([]*files.File{
		files.MustNewFileFromSource(files.NewBytesSource("tpl.yml", yamlTplData)),
		files.MustNewFileFromSource(files.NewBytesSource("values/data.yml", yamlData)),
	})

	ui := cmdcore.PlainUI{}
	opts := cmdtpl.NewOptions()


	out := opts.RunWithFiles(cmdtpl.TemplateInput{Files: filesToProcess}, ui)
	file := out.Files[0]

	return string(file.Bytes())
}

type jsFunc func(js.Value, []js.Value) interface{}

func registerFunc(name string, fn jsFunc) {
	js.Global().Set(name, js.FuncOf(fn))
	fmt.Printf("Registered \"%s\" with Global.\n", name)
}

func main() {
	registerFunc("add", add)
	registerFunc("template", doTemplate)
	<- make(chan int)
}
