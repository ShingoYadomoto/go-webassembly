package main

import (
	"syscall/js"
)

func main() {
	document := js.Global().Get("document")

	body := document.Call("getElementsByTagName", "body").Index(0)
	input := document.Call("getElementById", "colorCodeInput")

	cb := js.FuncOf(func(this js.Value, args []js.Value) interface{}{
		body.Get("style").Set("backgroundColor", input.Get("value").String())
		return nil
	})
	document.Call("getElementById", "colorChangeButton").Call("addEventListener", "click", cb)

	<-make(chan struct{}, 0)
}
