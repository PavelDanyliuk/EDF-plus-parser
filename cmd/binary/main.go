package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/PavelDanyliuk/EDF-plus-parser/services"
)

func main() {
	if err := run(); err != nil {
		fmt.Println("Fatal:", err)
	}
}

var filePath = flag.String("filePath", "./content/example_small.edf", "Path to the EDF file.")

func run() error {
	flag.Parse()

	file, err := os.Open(*filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := services.InitReader(file)
	fmt.Println(reader.Headers.GetHeadersJSON())

	return nil
}
