//+build js,wasm

package goes

import (
	"html/template"
	"syscall/js"
)

func LoadTemplate(temp *template.Template, id string) {
	temp = template.Must(
		temp.New(id).Parse(
			js.Global().Get("document").
				Call("getElementById", id).
				Get("innerHTML").String()),
	)
}
