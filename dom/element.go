//+build js,wasm

package dom

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"
	"syscall/js"
)

type Element struct {
	*js.Value
}

func NewElement(tag string, attributes map[string]string) *Element {
	v := js.Global().Get("document").Call("createElement", tag)
	el := &Element{Value: &v}
	for key, val := range attributes {
		el.SetAttribute(key, val)
	}
	return el
}

func NewElementFromTemplate(tmp *template.Template, name string, data interface{}) (*Element, error) {
	buf := bytes.NewBuffer(nil)
	err := tmp.ExecuteTemplate(buf, name, data)
	if err != nil {
		return nil, err
	}

	div := js.Global().Get("document").Call("createElement", "div")
	div.Set("innerHTML", strings.TrimSpace(buf.String()))

	v := div.Get("firstChild")
	if !v.Truthy() {
		return nil, fmt.Errorf("could not get created element")
	}

	return &Element{
		Value: &v,
	}, nil
}

func GetElementByID(id string) *Element {
	v := js.Global().Get("document").Call("getElementById", id)
	el := &Element{
		Value: &v,
	}
	if !el.Value.Truthy() {
		return nil
	}
	return el
}

func QuerySelector(query string) (*Element, error) {
	v := js.Global().Get("document").Call("querySelector", query)
	el := &Element{
		Value: &v,
	}

	if !el.Value.Truthy() {
		return nil, fmt.Errorf("query failed to find element matching selector: %q", query)
	}
	return el, nil
}

func QuerySelectorAll(query string) []*Element {
	var elements []*Element

	matches := js.Global().Get("document").Call("querySelectorAll", query)

	if !matches.Truthy() {
		return nil
	}

	length := matches.Length()

	for index := 0; index < length; index++ {
		v := matches.Index(index)

		el := &Element{
			Value: &v,
		}

		elements = append(elements, el)
	}

	return elements
}

func Body() *Element {
	bodyValue := js.Global().Get("document").Call("getElementsByTagName", "body").Index(0)
	return &Element{
		Value: &bodyValue,
	}
}

func (el *Element) Attribute(key string) string {
	return el.Call("getAttribute", key).String()
}

func (el *Element) SetAttribute(key, val string) {
	el.Call("setAttribute", key, val)
}

func (el *Element) ChildCount() int {
	return el.Get("childElementCount").Int()
}

func (el *Element) QuerySelector(query string) (*Element, error) {
	v := el.Call("querySelector", query)
	child := &Element{
		Value: &v,
	}
	if !child.Value.Truthy() {
		return nil, fmt.Errorf("query failed to find element matching selector: %q", query)
	}
	return child, nil
}

type EventListenerOptions struct {
	Capture, Once, Passive bool
}

func (el *Element) AddEventListener(eventName string, fn js.Func, options *EventListenerOptions) {
	params := []interface{}{js.ValueOf(eventName), fn}
	if options != nil {
		params = append(params, js.ValueOf(map[string]interface{}{
			"capture": options.Capture,
			"once":    options.Once,
			"passive": options.Passive,
		}))
	}
	el.Call("addEventListener", params...)
}

func (el *Element) Tag() string {
	return el.Get("tagName").String()
}

func (el *Element) JSValue() js.Value {
	if el == nil || el.Value == nil {
		return js.Null()
	}
	return *el.Value
}

func (el *Element) Remove() {
	if el == nil || el.Value == nil || el.Type() == js.TypeUndefined || el.Type() == js.TypeNull {
		return
	}
	el.Call("remove")
}
