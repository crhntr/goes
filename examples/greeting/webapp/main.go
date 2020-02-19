//+build js,wasm

package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"net/http"
	"syscall/js"

	"github.com/crhntr/goes"
	"github.com/crhntr/goes/dom"
	"github.com/crhntr/goes/examples/greeting"
)

type Page struct {
	Templates *template.Template
}

func main() {
	var page Page

	page.Templates = template.New("")
	goes.LoadTemplate(page.Templates, "page-template")

	greeting := js.ValueOf(map[string]interface{}{})
	greeting.Set("Reverse", js.FuncOf(page.Reverse))
	greeting.Set("FetchSpanish", js.FuncOf(page.FetchSpanish))
	js.Global().Set("greeting", greeting)

	<-make(chan struct{})
}

func (_ Page) Reverse(_ js.Value, args []js.Value) interface{} {
	msg := args[0].Get("innerHTML").String()
	msg = greeting.Reverse(msg)
	args[0].Set("innerHTML", msg)
	return nil
}

func (page *Page) FetchSpanish(_ js.Value, args []js.Value) interface{} {
	args[0].Set("disabled", true)
	go func() {

		req, _ := http.NewRequest(http.MethodGet, "/api/spanish-greeting", nil)
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			dom.Console.Logf("%q", err)
			return
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			dom.Console.Logf("%q", err)
			return
		}

		msg, err := ioutil.ReadAll(res.Body)
		if err != nil {
			dom.Console.Logf("%q", err)
			return
		}

		greeting := greeting.Envelope{Message: string(msg)}

		buf := bytes.NewBuffer([]byte(nil))
		if err := page.Templates.ExecuteTemplate(buf, "page-template", greeting); err != nil {
			dom.Console.Logf("%q", err)
			return
		}

		body, _ := dom.QuerySelector("body")

		body.Set("innerHTML", "")
		body.Set("innerHTML", buf.String())
	}()

	return nil
}
