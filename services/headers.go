package services

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/PavelDanyliuk/EDF-plus-parser/specification"
)

type Headers struct {
	Version            string   `json:"version"`
	PatientId          string   `json:"patiend_id"`
	RecordId           string   `json:"record_id"`
	StartDate          string   `json:"start_date"`
	StartTime          string   `json:"start_time"`
	HeaderSize         string   `json:"header_Size"`
	Reserved1          string   `json:"reserved1"`
	NumberOfRecords    string   `json:"number_of_records"`
	RecordDuration     string   `json:"record_duration"`
	NumberOfSignals    string   `json:"number_of_signals"`
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

func (h *Headers) Parse(reader io.Reader) {
	temp := make(map[string]string)

	/**
	 * Reading successively number of bytes and write them to the data.
	 * header_Size - represents how many bytes are reserved for a header.
	 */
	for _, header := range specification.HeadersStringLike {
		data := make([]byte, header.Size)

		_, err := io.ReadFull(reader, data)
		if err != nil {
			panic("Can not read a byte range") // TODO: handle errors properly
		}

		temp[header.Name] = strings.TrimSpace(string(data))
	}

	h.Version = temp["version"]
	h.PatientId = temp["patient_id"]
	h.RecordId = temp["record_id"]
	h.StartDate = temp["start_date"]
	h.StartTime = temp["start_time"]
	h.HeaderSize = temp["header_Size"]
	h.Reserved1 = temp["reserved1"]
	h.NumberOfRecords = temp["number_of_records"]
	h.RecordDuration = temp["record_duration"]
	h.NumberOfSignals = temp["number_of_signals"]

	temp2 := make(map[string][]string)
	for _, header := range specification.HeadersArrayLike {
		var result []string
		signals, _ := strconv.Atoi(h.NumberOfSignals)

		for i := 0; i < signals; i++ {
			data := make([]byte, header.Size)
			_, err := io.ReadFull(reader, data)
			if err != nil {
				panic("Can not read a byte range") // TODO: handle errors properly
			}

			result = append(result, strings.TrimSpace(string(data)))
		}

		temp2[header.Name] = result
	}

	h.Labels = temp2["labels"]
	h.TransducerTypes = temp2["transducer_types"]
	h.PhysicalDimensions = temp2["physical_dimensions"]
	h.PhysicalMinimums = temp2["physical_minimums"]
	h.PhysicalMaximums = temp2["physical_maximums"]
	h.DigitalMinimums = temp2["digital_minimums"]
	h.DigitalMaximums = temp2["digital_maximums"]
	h.Prefiltering = temp2["prefiltering"]
	h.SamplesPerRecord = temp2["samples_per_record"]
	h.Reserved2 = temp2["reserved2"]
}

func (h *Headers) GetHeadersJSON() ([]byte, error) {
	json, err := json.Marshal(h)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return json, nil
}
