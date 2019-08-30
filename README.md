# GOES

Write the WASM test!

![Go write the WASM test (a riff on Bernie's quote "I wrote the damn bill!")](https://raw.githubusercontent.com/crhntr/goes/master/tooling/GoWriteTheWasmTest.png)

Inspired by [go-billy](https://github.com/src-d/go-billy), "goes" wraps
"syscall/js" with interface types. The goal is to make it easier to test the Go
side of interactions with the Javascript runtime without needing to run the
project in a browser or with Node.

The example in examples/greeting has an acceptance test using
`github.com/chromedp/chromedp` to test how the WASM interacts with the page.

## Contributing

- Please cake sure to run `go generate ./...` when contributing to the interface declarations.

- To run acceptance tests for the examples run `go test -tags acceptance ./...`.

## Examples

Given a this wrapper around a div containing a greeting,

```go
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
```

it can be tested as follows:

*Note the build flag ensures the unit tests are not run in the browser
when the acceptance tests are being run.*

```go
//+build !js

package greeting_test

import (
	"github.com/crhntr/goes/examples/greeting"
	"github.com/crhntr/goes/goesfakes"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestNewHelloBox(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	div := goesfakes.NewValue(ctrl)
	div.EXPECT().Call(gomock.Eq("setAttribute"), gomock.Eq("id"), gomock.Eq("hello-box"))
	div.EXPECT().Set(gomock.Eq("innerText"), gomock.Eq("hello"))

	document := goesfakes.NewValue(ctrl)
	document.EXPECT().Call(gomock.Eq("createElement"), gomock.Eq("div")).Return(div).Times(1)

	box := greeting.NewHelloBox("hello")

	box.Create(document)
}

func TestHelloBox_SetMessage(t *testing.T) {
	{
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		div := goesfakes.NewValue(ctrl)
		gomock.InOrder(
			div.EXPECT().Call(gomock.Eq("setAttribute"), gomock.Eq("id"), gomock.Eq("hello-box")),

			div.EXPECT().Set(gomock.Eq("innerText"), gomock.Eq("hello")),
			div.EXPECT().Set(gomock.Eq("innerText"), gomock.Eq("Hello, world!")),
		)

		document := goesfakes.NewValue(ctrl)
		document.EXPECT().Call(gomock.Eq("createElement"), gomock.Eq("div")).Return(div).Times(1)

		box := greeting.NewHelloBox("hello")

		box.Create(document)
		box.SetMessage("Hello, world!")
		box.Message()
	}

	t.Run("when initialized with an empty string", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		msgTxtFromDom := goesfakes.NewValue(ctrl)
		msgTxtFromDom.EXPECT().String().Return("Hello, world!").AnyTimes()

		div := goesfakes.NewValue(ctrl)
		gomock.InOrder(
			div.EXPECT().Call(gomock.Eq("setAttribute"), gomock.Eq("id"), gomock.Eq("hello-box")),
			div.EXPECT().Set(gomock.Eq("innerText"), gomock.Eq("")),

			div.EXPECT().Get(gomock.Eq("innerText")).Return(msgTxtFromDom), // <- what changed
		)

		document := goesfakes.NewValue(ctrl)
		document.EXPECT().Call(gomock.Eq("createElement"), gomock.Eq("div")).Return(div).Times(1)

		box := greeting.NewHelloBox("")

		box.Create(document)
		box.Message()
	})
}

func TestHelloBox_SetMessagz(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	msgTxtFromDom := goesfakes.NewValue(ctrl)
	msgTxtFromDom.EXPECT().String().Return("Hello, world!").AnyTimes()

	div := goesfakes.NewValue(ctrl)
	gomock.InOrder(
		div.EXPECT().Call(gomock.Eq("setAttribute"), gomock.Eq("id"), gomock.Eq("hello-box")),
		div.EXPECT().Set(gomock.Eq("innerText"), gomock.Eq("hello")),
		div.EXPECT().Set(gomock.Eq("innerText"), gomock.Eq("Hello, world!")),
		// div.EXPECT().Get(gomock.Eq("innerText")).Return(msgTxtFromDom),
	)

	document := goesfakes.NewValue(ctrl)
	document.EXPECT().Call(gomock.Eq("createElement"), gomock.Eq("div")).Return(div).Times(1)

	box := greeting.NewHelloBox("hello")

	box.Create(document)
	box.SetMessage("Hello, world!")
	box.Message()
}
```
