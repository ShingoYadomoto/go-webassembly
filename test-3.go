package main

import (
	"fmt"
	"image"
	"syscall/js"

	"bytes"
	"encoding/base64"
	"image/color"
	"image/png"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/anthonynsimon/bild/adjust"
	"github.com/anthonynsimon/bild/imgio"
)

func main() {
	filename := js.Global().Get("document").
		Call("getElementById", "img").
		Get("src").String()

	response, err := http.Get(filename)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	img, err := png.Decode(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	bounds := img.Bounds()
	dest := image.NewGray16(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := color.Gray16Model.Convert(img.At(x, y))
			gray, _ := c.(color.Gray16)
			dest.Set(x, y, gray)
		}
	}

	fmt.Println("5")
	f := &os.File{}

	err = png.Encode(f, dest)

	fmt.Println("6")
	data := make([]byte, 1000000000000000000)
	f.Read(data)
	fmt.Println("7")

	fmt.Println(base64.StdEncoding.EncodeToString(data))

	cb := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// quick return if no source image is yet uploaded
		if sh.sourceImg == nil {
			return nil
		}
		delta := args[0].Get("img").Get("valueAsNumber").Int()
		start := time.Now()
		res := adjust.Hue(sh.sourceImg, delta)
		sh.updateImage(res, start)
		args[0].Call("preventDefault")
		return nil
	})

	js.Global().Get("document").
		Call("getElementById", "hue").
		Call("addEventListener", "mousemove", cb)

	<-make(chan struct{}, 0)
}

//
//type Shimmer struct {
//	inBuf                    []uint8
//	outBuf                   bytes.Buffer
//	onImgLoadCb, shutdownCb  js.Func
//	brightnessCb, contrastCb js.Func
//	hueCb, satCb             js.Func
//	sourceImg                image.Image
//}
//
//// updateImage writes the image to a byte buffer and then converts it to base64.
//// Then it sets the value to the src attribute of the target image.
func updateImage(img *image.RGBA, start time.Time) {
	buf := &bytes.Buffer{}
	enc := imgio.PNGEncoder()
	err := enc(buf, img)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	dst := js.Global().Get("Uint8Array").New(len(buf.Bytes()))
	n := js.CopyBytesToJS(dst, buf.Bytes())
	fmt.Println(err.Error())

	fmt.Println("bytes copied:", strconv.Itoa(n))
	js.Global().Call("displayImage", dst)
	fmt.Println("time taken:", time.Now().Sub(start).String())
	buf.Reset()
}

//
//// 画像ロード
//func (s *Shimmer) setupOnImgLoadCb() {
//	s.onImgLoadCb = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
//		array := args[0]
//		s.inBuf = make([]uint8, array.Get("byteLength").Int())
//		js.CopyBytesToGo(s.inBuf, array)
//
//		reader := bytes.NewReader(s.inBuf)
//		var err error
//		s.sourceImg, _, err = image.Decode(reader)
//		if err != nil {
//			fmt.Println(err.Error())
//			return nil
//		}
//		fmt.Println("Ready for operations")
//
//		return nil
//	})
//}
