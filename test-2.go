package main

import (
	"syscall/js"
	"math/rand"
	"time"
)

var (
	img = js.Global().Call("eval", "new Image()")
	document = js.Global().Get("document")
)

func main() {
	img.Set("src", "./image/gopher.png")

	body := document.Call("getElementsByTagName", "body").Index(0)
	body.Call("addEventListener", "mousemove", js.FuncOf(func(this js.Value, args []js.Value) interface{}{
		render()
		return nil
	}))

	<-make(chan struct{}, 0)
}


func render(){
	canvas := document.Call("getElementById", "sample")
	ctx := canvas.Call("getContext", "2d")

	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(1000)
	rand.Seed(time.Now().UnixNano()+100000)
	y:= rand.Intn(1000)

	ctx.Call("drawImage", img, x,y)
}

