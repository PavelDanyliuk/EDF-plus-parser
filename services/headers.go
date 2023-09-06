package services

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/PavelDanyliuk/EDF-plus-parser/specification"
)

type Headers struct {
	file *[]byte

	Version            string   `json:"version"`
	PatientId          string   `json:"patiend_id"`
	RecordId           string   `json:"record_id"`
	StartDate          string   `json:"start_date"`
	StartTime          string   `json:"start_time"`
	HeaderSize         string   `json:"header_Size"`
	Reserved1          string   `json:"reserved1"`
	NumberOfRecords    string   `json:"number_of_records"`
	RecordDuration     string   `json:"record_duration"`
	NumberOfSignals    int      `json:"number_of_signals"`
	Labels             []string `json:"labels"`
	TransducerTypes    []string `json:"transducer_types"`
	PhysicalDimensions []string `json:"physical_dimensions"`
	PhysicalMinimums   []string `json:"physical_minimums"`
	PhysicalMaximums   []string `json:"physical_maximums"`
	DigitalMinimums    []string `json:"digital_minimums"`
	DigitalMaximums    []string `json:"digital_maximums"`
	Prefiltering       []string `json:"prefiltering"`
	SamplesPerRecord   []string `json:"smples_per_record"`
	Reserved2          []string `json:"reserved2"`
}

func (h *Headers) ParseHeaders() {
	temp := make(map[string][]string)

	signals, _ := strconv.Atoi(
		strings.TrimSpace(
			h.readHeader("number_of_signals")[0],
		),
	) // []string -> int

	h.NumberOfSignals = signals

	for key := range specification.HeadersSpecification {
		temp[key] = h.readHeader(key)
	}

	h.Version = temp["version"][0]
	h.PatientId = temp["patient_id"][0]
	h.RecordId = temp["record_id"][0]
	h.StartDate = temp["start_date"][0]
	h.StartTime = temp["start_time"][0]
	h.HeaderSize = temp["header_Size"][0]
	h.Reserved1 = temp["reserved1"][0]
	h.NumberOfRecords = temp["number_of_records"][0]
	h.RecordDuration = temp["record_duration"][0]
	h.Labels = temp["labels"]
	h.TransducerTypes = temp["transducer_types"]
	h.PhysicalDimensions = temp["physical_dimensions"]
	h.PhysicalMinimums = temp["physical_minimums"]
	h.PhysicalMaximums = temp["physical_maximums"]
	h.DigitalMinimums = temp["digital_minimums"]
	h.DigitalMaximums = temp["digital_maximums"]
	h.Prefiltering = temp["prefiltering"]
	h.SamplesPerRecord = temp["samples_per_record"]
	h.Reserved2 = temp["reserved2"]
}

func (h *Headers) GetHeadersJSON() ([]byte, error) {
	json, err := json.Marshal(h)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return json, nil
}

func (h *Headers) readHeader(headerName string) []string {
	target := specification.HeadersSpecification[headerName]
	start, end := h.byteRange(headerName)
	header := (*h.file)[start : end+1]

	if target.IsArray {
		var result []string

		for i := 0; i < h.NumberOfSignals; i++ {
			val := header[i*target.Size : (i+1)*target.Size]
			// string() converts byte to string
			// .Fields splits a string by words
			result = append(result, strings.Fields(strings.TrimSpace(string(val)))...)
		}

		return result
	}

	return []string{strings.TrimSpace(string(header))}
}

func (h *Headers) byteRange(headerName string) (int, int) {
	target, ok := specification.HeadersSpecification[headerName]

	if !ok {
		panic("Can not read a header name from the specification!" + "Header: " + headerName)
	}

	startByte := 0
	endByte := 0

	for i := 0; i < target.Position; i++ {
		header := specification.HeadersSpecification[specification.HeadersOrder[i]]

		if header.IsArray {
			startByte += h.NumberOfSignals * header.Size
		} else {
			startByte += header.Size
		}
	}

	targetSize := 0

	if target.IsArray {
		targetSize = h.NumberOfSignals * target.Size
	} else {
		targetSize = target.Size
	}

	endByte = startByte + targetSize - 1

	return startByte, endByte
}
