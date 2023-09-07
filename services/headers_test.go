package services

import (
	"os"
	"testing"
)

const filePath string = "../content/example_small.edf"

func TestInitReader(t *testing.T) {
	file, err := os.Open(filePath)
	if err != nil {
		t.Errorf("Can not read the file")
	}
	defer file.Close()

	reader := InitReader(file)

	expectedSignals := "5"
	if reader.Headers.NumberOfSignals != expectedSignals {
		t.Errorf("Reader.signals = %s, expected %s", reader.Headers.NumberOfSignals, expectedSignals)
	}
}
