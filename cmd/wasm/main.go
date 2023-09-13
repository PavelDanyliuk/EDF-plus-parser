//go:build !test

package main

import (
	"syscall/js"

	"github.com/PavelDanyliuk/EDF-plus-parser/services"
)

var reader *services.Reader

func main() {
	channel := make(chan struct{}, 0)

	js.Global().Set("initReader", js.FuncOf(initReader))
	js.Global().Set("getHeaders", js.FuncOf(getHeaders))
	js.Global().Set("getSignals", js.FuncOf(getSignals))

	<-channel

}

func initReader(this js.Value, args []js.Value) interface{} {
	byteSlice := make([]byte, args[0].Length())
	js.CopyBytesToGo(byteSlice, args[0])

	reader = services.InitReader(byteSlice)
	return nil
}

func getHeaders(this js.Value, args []js.Value) interface{} {
	json, _ := reader.Headers.GetHeadersJSON()
	jsData := js.Global().Get("JSON").Call("parse", string(json))

	return jsData
}

func getSignals(this js.Value, args []js.Value) interface{} {
	json, _ := reader.Channels.GetAllSignalsJSON()
	jsData := js.Global().Get("JSON").Call("parse", string(json))

	return jsData
}
