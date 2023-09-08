package services

import (
	"bufio"
	"os"
	"reflect"
	"testing"
)

const filePath string = "../content/example_small.edf"

func TestParseHeaders(t *testing.T) {
	file, err := os.Open(filePath)
	if err != nil {
		t.Errorf("Can not read the file")
	}

	defer file.Close()

	bufioReader := bufio.NewReader(file)
	headers := &Headers{}

	headers.Parse(bufioReader)

	expectedSignals := "5"
	if headers.NumberOfSignals != expectedSignals {
		t.Errorf("headers.NumberOfSignals = %s, expected %s", headers.NumberOfSignals, expectedSignals)
	}

	expectedLabels := []string{"SPO2", "Heart Rate", "Effort", "Flow", "EDF Annotations"}
	if !reflect.DeepEqual(headers.Labels, expectedLabels) {
		t.Errorf("headers.Labels = %s, expected %s", headers.Labels, expectedLabels)
	}
}
