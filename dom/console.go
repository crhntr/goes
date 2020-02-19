//+build js,wasm

package dom

import (
	"fmt"
	"syscall/js"
)

type console string

const Console console = "console"

func (_ console) Logf(format string, args ...interface{}) {
	js.Global().Get("console").Call("log", fmt.Sprintf(format, args...))
}
