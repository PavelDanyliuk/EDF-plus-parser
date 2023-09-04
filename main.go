package main

import (
	"fmt"
	"os"

	"github.com/PavelDanyliuk/EDF-plus-parser/services"
)

func main() {
	if err := run(); err != nil {
		fmt.Println("Fatal:", err)
	}
}

const filePath string = "./content/example_small.edf"

func run() error {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	reader := services.InitReader(&file)
	headers := reader.GetHeaders()

	for key, value := range headers {
		fmt.Println(key, "-", value)
	}

	return nil
}
