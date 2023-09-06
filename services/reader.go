package services

type Reader struct {
	Headers     *Headers
	Channels    *Channels
	Annotations *Annotations
}

func InitReader(file *[]byte) *Reader {
	headers := &Headers{file: file}
	channels := &Channels{file: file}
	annotations := &Annotations{file: file}

	reader := &Reader{
		Headers:     headers,
		Channels:    channels,
		Annotations: annotations,
	}

	reader.Headers.ParseHeaders()

	return reader
}
