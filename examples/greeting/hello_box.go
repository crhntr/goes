package greeting

import (
	"github.com/crhntr/goes"
)

type HelloBox struct {
	div     goes.Value
	message string
}

func NewHelloBox(msg string) *HelloBox {
	return &HelloBox{
		message: msg,
	}
}

func (box *HelloBox) SetMessage(msg string) {
	box.message = msg
	box.div.Set("innerText", msg)
}

func (box *HelloBox) Message() string {
	if box.message != "" {
		return box.message
	}
	box.message = box.div.Get("innerText").String()
	return box.message
}

func (box *HelloBox) Create(document goes.Value) goes.Value {
	box.div = document.Call("createElement", "div")
	box.div.Call("setAttribute", "id", "hello-box")
	box.SetMessage(box.message)
	return box.div
}
