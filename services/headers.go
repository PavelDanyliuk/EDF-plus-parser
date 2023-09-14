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
	NumberOfRecords    int      `json:"number_of_records"`
	RecordDuration     int      `json:"record_duration"`
	NumberOfSignals    int      `json:"number_of_signals"`
	PhysicalDimensions []int    `json:"physical_dimensions"`
	PhysicalMinimums   []int    `json:"physical_minimums"`
	PhysicalMaximums   []int    `json:"physical_maximums"`
	DigitalMinimums    []int    `json:"digital_minimums"`
	DigitalMaximums    []int    `json:"digital_maximums"`
	SamplesPerRecord   []int    `json:"smples_per_record"`
	Labels             []string `json:"labels"`
	TransducerTypes    []string `json:"transducer_types"`
	Prefiltering       []string `json:"prefiltering"`
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

	numberOfRecords, _ := strconv.Atoi(temp["number_of_records"])
	h.NumberOfRecords = numberOfRecords

	recordDuration, _ := strconv.Atoi(temp["record_duration"])
	h.RecordDuration = recordDuration

	numberOfSignals, _ := strconv.Atoi(temp["number_of_signals"])
	h.NumberOfSignals = numberOfSignals

	temp2 := make(map[string][]string)
	for _, header := range specification.HeadersArrayLike {
		var result []string

		for i := 0; i < h.NumberOfSignals; i++ {
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
	h.Prefiltering = temp2["prefiltering"]
	h.Reserved2 = temp2["reserved2"]
	h.PhysicalDimensions = convertToInt(temp2["physical_dimensions"])
	h.PhysicalMinimums = convertToInt(temp2["physical_minimums"])
	h.PhysicalMaximums = convertToInt(temp2["physical_maximums"])
	h.DigitalMinimums = convertToInt(temp2["digital_minimums"])
	h.DigitalMaximums = convertToInt(temp2["digital_maximums"])
	h.SamplesPerRecord = convertToInt(temp2["samples_per_record"])
}

func (h *Headers) GetHeadersJSON() ([]byte, error) {
	json, err := json.Marshal(h)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return json, nil
}

func convertToInt(input []string) []int {
	result := make([]int, len(input))

	for i, str := range input {
		val, _ := strconv.Atoi(str)
		result[i] = val
	}

	return result
}
