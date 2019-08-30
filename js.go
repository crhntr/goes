// +build js

package goes

import (
	"syscall/js"
)

type JSValuer interface {
	Value

	JSValue() js.Value
}

type Wrapper = js.Wrapper

func (v V) JSValue() js.Value { return v.Value }

type Runtime struct{}

func (_ Runtime) Global() Value { return V{js.Global()} }

func (_ Runtime) Null() Value { return V{js.Null()} }

func (_ Runtime) Undefined() Value { return V{js.Undefined()} }

func (_ Runtime) ValueOf(x interface{}) Value { return V{js.ValueOf(x)} }

type V struct{ js.Value }

func (v V) Bool() bool                               { return v.Value.Bool() }
func (v V) Call(m string, args ...interface{}) Value { return V{v.Value.Call(m, args...)} }
func (v V) Float() float64                           { return v.Value.Float() }
func (v V) Get(p string) Value                       { return V{v.Value.Get(p)} }
func (v V) Index(i int) Value                        { return V{v.Value.Index(i)} }
func (v V) InstanceOf(t Value) bool {
	val, ok := t.(V)
	if !ok {
		panic("could not determine type of Value in method InstanceOf")
	}
	return v.Value.InstanceOf(val.Value)
}
func (v V) Int() int                         { return v.Value.Int() }
func (v V) Invoke(args ...interface{}) Value { return V{v.Value.Invoke(args...)} }
func (v V) Length() int                      { return v.Value.Length() }
func (v V) New(args ...interface{}) Value    { return V{v.Value.New(args...)} }
func (v V) Set(p string, x interface{})      { v.Value.Set(p, x) }
func (v V) SetIndex(i int, x interface{})    { v.Value.SetIndex(i, x) }
func (v V) String() string                   { return v.Value.String() }
func (v V) Truthy() bool                     { return v.Value.Truthy() }
func (v V) Type() Type                       { return Type(v.Value.Type()) }

func FuncOf(fn Function) Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var interfaceArgs []Value

		for _, arg := range args {
			interfaceArgs = append(interfaceArgs, V{arg})
		}

		return fn(V{this}, interfaceArgs)
	})
}
