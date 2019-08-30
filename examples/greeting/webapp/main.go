//+build js,wasm

package main

import (
	"io/ioutil"
	"net/http"

	"github.com/crhntr/goes"
	"github.com/crhntr/goes/examples/greeting"
)

type Console struct {
	goes.Value
}

func (console Console) Log(args ...interface{}) {
	console.Call("log", args...)
}

func main() {
	js := goes.Runtime{}
	console := Console{js.Global().Get("console")}
	console.Log("Hello, world!")

	document := js.Global().Get("document")

	box := greeting.NewHelloBox("Hello, world!")
	parent := document.Call("getElementById", "main")
	parent.Call("appendChild", box.Create(document))

	req, _ := http.NewRequest(http.MethodGet, "/api/spanish-greeting", nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		console.Log(err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		console.Log(res.StatusCode)
		return
	}
	msg, err := ioutil.ReadAll(res.Body)
	if err != nil {
		console.Log(err)
		return
	}

	box.SetMessage(string(msg))

	<-make(chan struct{})
}
