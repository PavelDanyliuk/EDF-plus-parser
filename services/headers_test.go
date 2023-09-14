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
	channels := &Channels{
		headers: headers,
	}

	headers.Parse(bufioReader)
	channels.Parse(bufioReader)

	expectedSignals := 5
	if headers.NumberOfSignals != expectedSignals {
		t.Errorf("headers.NumberOfSignals = %d, expected %d", headers.NumberOfSignals, expectedSignals)
	}

	expectedLabels := []string{"SPO2", "Heart Rate", "Effort", "Flow", "EDF Annotations"}
	if !reflect.DeepEqual(headers.Labels, expectedLabels) {
		t.Errorf("headers.Labels = %s, expected %s", headers.Labels, expectedLabels)
	}

	expectedSignalsData := []float64{90.4974441138323, 93.00007629510948, 90.4974441138323}
	actualSignalsData := channels.data[0][:3]
	if !reflect.DeepEqual(actualSignalsData, expectedSignalsData) {
		t.Errorf("Channels data doesn't parse correctly \n Expected: %b \n Actual: %b", expectedSignalsData, actualSignalsData)
	}
}
