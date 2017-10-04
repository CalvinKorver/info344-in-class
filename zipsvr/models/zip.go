package models

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

// Zip ...
type Zip struct {
	Code  string
	City  string
	State string
}

// ZipSlice ...
// 64 bit number that points to a memory address
type ZipSlice []*Zip

// ZipIndex ...
type ZipIndex map[string]ZipSlice

// LoadZips ...
func LoadZips(fileName string) (ZipSlice, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}

	reader := csv.NewReader(f)
	_, err = reader.Read() // Underscore means ignore
	if err != nil {
		return nil, fmt.Errorf("error reading header row: %v", err)
	}

	// zips := ZipSlice{} This would just allocate a 0 capactiy, which we don't want

	zips := make(ZipSlice, 0, 43000) // Efficient because we pre-allocate the memory
	for {
		fields, err := reader.Read()
		if err == io.EOF {
			return zips, nil
		}

		if err != nil {
			return nil, fmt.Errorf("error reading record: %v", err)
		}
		z := &Zip{
			Code:  fields[0],
			City:  fields[3],
			State: fields[6],
		}
		zips = append(zips, z) // Appends existing slice and returns it
	}
}
