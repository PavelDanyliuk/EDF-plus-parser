package services

import (
	"encoding/binary"
	"encoding/json"
	"io"
)

type Channels struct {
	data    [][]float64
	headers *Headers
}

func (c *Channels) Parse(input io.Reader) {
	nr := c.headers.NumberOfRecords
	ns := c.headers.NumberOfSignals
	c.data = make([][]float64, ns)

	for i := 0; i < ns; i++ {
		c.data[i] = make([]float64, nr*c.headers.SamplesPerRecord[i])
	}

	for i := 0; i < nr; i++ {
		for j := 0; j < ns; j++ {

			sampels := c.headers.SamplesPerRecord[j]
			for s := 0; s < sampels; s++ {
				var rawValue int16
				err := binary.Read(input, binary.LittleEndian, &rawValue)
				if err != nil {
					panic(err)
				}

				scaledValue := c.scaleToPhysical(rawValue, j)
				c.data[j][i*sampels+s] = scaledValue
			}
		}
	}
}

func (c *Channels) GetAllSignalsJSON() ([]byte, error) {
	json, err := json.Marshal(c.data)
	if err != nil {
		return nil, err
	}

	return json, nil
}

func (c *Channels) scaleToPhysical(value int16, channelIndex int) float64 {
	dMin := float64(c.headers.DigitalMinimums[channelIndex])
	dMax := float64(c.headers.DigitalMaximums[channelIndex])
	pMin := float64(c.headers.PhysicalMinimums[channelIndex])
	pMax := float64(c.headers.PhysicalMaximums[channelIndex])

	a := (pMax - pMin) / (dMax - dMin)
	b := pMin - a*dMin

	return float64(value)*a + b
}
