package services

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

type Reader struct {
	Headers     *Headers
	Channels    *Channels
	Annotations *Annotations
}

func InitReader(source interface{}) *Reader {
	var r io.Reader

	switch src := source.(type) {
	case []byte:
		r = bytes.NewReader(src)
	case *os.File:
		r = bufio.NewReader(src)
	}

	headers := &Headers{}
	channels := &Channels{
		headers: headers,
	}
	annotations := &Annotations{}

	reader := &Reader{
		Headers:     headers,
		Channels:    channels,
		Annotations: annotations,
	}

	reader.Headers.Parse(r)
	reader.Channels.Parse(r)

	return reader
}
