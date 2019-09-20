//+build js,wasm

package main

import (
	"fmt"
	"strings"

	"github.com/crhntr/goes"
)

func main() {
	js := goes.Runtime{}

	// console := Console{js.Global().Get("console")}

	document := js.Global().Get("document")
	body := document.Get("body")

	if err := mount(body); err != nil {
    fmt.Println(err)
  }

	//
	// window := js.Global().Get("window")
	// console := Console{js.Global().Get("console")}
	//
	//

	<-make(chan struct{})
}

type Console struct {
	goes.Value
}

func (console Console) Log(args ...interface{}) {
	console.Call("log", args...)
}

func mount(val goes.Value) error {
  isElement := func(v goes.Value) bool {
    if n.Type() != goes.TypeObject {
      return false
    }
    if t := n.Get("nodeType"); t.Int() != 1 {
      return false
    }
    return true
  }

  if !isElement(val) {
    return errors.New("mount node is not en element")
  }

  childNodes := val.Get("childNodes")
  childNodeLength := childNodes.Get("length").Int()
  for index := 0; index < childNodeLength; index++ {
    node := childNodes.Index(index)

    ats := node.Get("attributes")
    if ats.Type() != goes.TypeObject {
      continue
    }

    atsLength := ats.Get("length")
    for indexAt := 0; indexAt < atsLength.Int(); indexAt++ {
      at := ats.Index(indexAt)
      if at.Type() == goes.TypeUndefined {
        continue
      }
      name := at.Get("name").String()
      if strings.HasPrefix(name, "goes-") || strings.HasPrefix(name, "@") || strings.HasPrefix(name, ":") {
        fmt.Println(name)
      }
    }

    cns := node.Get("childNodes")
    // console.Log(cns)

    cnsl := cns.Get("length").Int()
    for i := 0; i < cnsl; i++ {
      n := cns.Index(i)

      if err := mount(n); err != nil {
        return err
      }
    }
  }
}


// console.Log(js.Global().Get("location").Get("pathname"))
// window.Call("addEventListener", "hashchange", goes.FuncOf(func(this goes.Value, args []goes.Value) interface{} {
// 	console.Log(js.Global().Get("location").Get("pathname"))
// 	return nil
// }), true)
