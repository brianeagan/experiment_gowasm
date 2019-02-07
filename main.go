package main

import (
	"bytes"
	"html/template"
	"log"
	"syscall/js"
	"github.com/gobuffalo/packr"
	"github.com/davecgh/go-spew/spew"
)


var templateBox packr.Box
var data = Content{
PageTitle: "WEBASMTITLE",
	Histories: []History{
	{Item: "one"}, {Item: "two"}, {Item: "three"},
	},
}
type History struct {
	Item string
}

type Content struct{
	PageTitle string
	Histories  [] History

}

//would use something to pack the templates into the binary
func renderThings(i []js.Value) {
	//tmpl string, vars struct{}
	spew.Dump(i)
	var rendered string
	html, err :=  templateBox.FindString(i[0].String())
	if err !=nil {
		log.Printf("Not a valid template %v", err)
		return
	}
	target := template.New(html)
	renderedBuf := new(bytes.Buffer)
	//would fetch and parse from backend
	target.Execute(renderedBuf, data)
	rendered = renderedBuf.String()

	//get the template and shove in stuff
	//rendered := target.Execute(io, vars)
	js.Global().Get("document").Call("getElementById", "target").Set("value", rendered)
}
//
//func add(i []js.Value) {
//	js.Global().Set("output", js.ValueOf(i[0].Int()+i[1].Int()))
//	println(js.ValueOf(i[0].Int() + i[1].Int()).String())
//}
//
//func subtract(i []js.Value) {
//	js.Global().Set("output", js.ValueOf(i[0].Int()-i[1].Int()))
//	println(js.ValueOf(i[0].Int() - i[1].Int()).String())
//}

func registerCallbacks() {
	//js.Global().Set("add", js.NewCallback(add))
	//js.Global().Set("subtract", js.NewCallback(subtract))
	js.Global().Set("render", js.NewCallback(renderThings))
}


func init()  {
	templateBox = packr.NewBox("./templates")
}

func main() {
	c := make(chan struct{}, 0)

	//pretend json was got here




	println("WASM Go Initialized")
	// register functions
	registerCallbacks()
	<-c
}