package services

import (
	"strconv"
	"strings"

	"github.com/PavelDanyliuk/EDF-plus-parser/specification"
)

type Reader struct {
	file    *[]byte
	signals int
}

func InitReader(file *[]byte) *Reader {
	r := &Reader{}
	r.file = file

	signals, _ := strconv.Atoi(
		strings.TrimSpace(
			r.readHeader("number_of_signals")[0],
		),
	) // []string -> int8

	r.signals = signals

	return r
}

func (r *Reader) GetHeaders() specification.ParsedHeaders {
	result := make(specification.ParsedHeaders)

	for _, v := range specification.HeadersOrder {
		result[v] = r.readHeader(v)
	}

	return result
}

func (r *Reader) readHeader(headerName string) []string {
	target := specification.HeadersSpecification[headerName]
	start, end := r.byteRange(headerName)
	header := (*r.file)[start : end+1]

	if target.IsArray {
		var result []string

		for i := 0; i < r.signals; i++ {
			val := header[i*target.Size : (i+1)*target.Size]
			// string() converts byte to string
			// .Fields splits a string by words
			result = append(result, strings.Fields(string(val))...)
		}

		return result
	}

	return []string{string(header)}
}

func (r *Reader) byteRange(headerName string) (int, int) {
	target, ok := specification.HeadersSpecification[headerName]

	if !ok {
		panic("Can not read a header name from the specification!" + "Header: " + headerName)
	}

	startByte := 0
	endByte := 0

	for i := 0; i < target.Position; i++ {
		header := specification.HeadersSpecification[specification.HeadersOrder[i]]

		if header.IsArray {
			startByte += r.signals * header.Size
		} else {
			startByte += header.Size
		}
	}

	targetSize := 0

	if target.IsArray {
		targetSize = r.signals * target.Size
	} else {
		targetSize = target.Size
	}

	endByte = startByte + targetSize - 1

	return startByte, endByte
}
