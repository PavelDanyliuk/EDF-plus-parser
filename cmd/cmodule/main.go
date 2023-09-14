//go:build !test

package main

import (
	"C"
	"log"
	"os"

	"github.com/PavelDanyliuk/EDF-plus-parser/services"
)

var reader *services.Reader

func main() {}

//export initReader
func initReader(filePath *C.char) {
	goFilePath := C.GoString(filePath)

	file, err := os.Open(goFilePath)
	if err != nil {
		log.Println("OPEN ERROR") // TODO: handle errors
	}
	defer file.Close()

	reader = services.InitReader(file)
}

//export getHeaders
func getHeaders() *C.char {
	headers, _ := reader.Headers.GetHeadersJSON() // TODO: do I need to handle error here?
	return C.CString(string(headers))
}

//export getChannels
func getChannels() *C.char {
	channels, _ := reader.Channels.GetAllSignalsJSON() // TODO: do I need to handle error here?
	return C.CString(string(channels))
}
