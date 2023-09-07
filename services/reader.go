package services

import (
	"bufio"
	"os"
)

type Reader struct {
	Headers     *Headers
	Channels    *Channels
	Annotations *Annotations
}

func InitReader(file *os.File) *Reader {
	bufioReader := bufio.NewReader(file)

	headers := &Headers{}
	channels := &Channels{}
	annotations := &Annotations{}

	reader := &Reader{
		Headers:     headers,
		Channels:    channels,
		Annotations: annotations,
	}

	reader.Headers.Parse(bufioReader)

	return reader
}
