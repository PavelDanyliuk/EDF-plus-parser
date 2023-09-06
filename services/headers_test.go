package services

import (
	"os"
	"testing"
)

const filePath string = "../content/example_small.edf"

func TestInitReader(t *testing.T) {
	file, err := os.ReadFile(filePath)

	if err != nil {
		t.Errorf("Can not read the file")
	}

	reader := InitReader(&file)

	expectedSignals := 5
	if reader.Headers.NumberOfSignals != expectedSignals {
		t.Errorf("Reader.signals = %d, expected %d", reader.Headers.NumberOfSignals, expectedSignals)
	}
}
