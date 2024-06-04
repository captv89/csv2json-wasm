package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"syscall/js"
)

// csv2json function to convert CSV content to JSON
func csv2json(csvContent string) (string, error) {
	reader := csv.NewReader(strings.NewReader(csvContent))
	headers, err := reader.Read()
	if err != nil {
		return "", fmt.Errorf("error reading header row: %w", err)
	}

	var records []map[string]string
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", fmt.Errorf("error reading row: %w", err)
		}

		record := make(map[string]string)
		for i, value := range row {
			record[headers[i]] = value
		}
		records = append(records, record)
	}

	jsonData, err := json.MarshalIndent(records, "", "    ")
	if err != nil {
		return "", fmt.Errorf("error converting to JSON: %w", err)
	}

	return string(jsonData), nil
}

// Exported function for WebAssembly
func Csv2Json(this js.Value, p []js.Value) interface{} {
	csvContent := p[0].String()
	jsonData, err := csv2json(csvContent)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}
	}
	return map[string]interface{}{
		"json": jsonData,
	}
}

func main() {
	fmt.Println("Go WebAssembly CSV to JSON Converter")

	// Register the function
	js.Global().Set("csv2json", js.FuncOf(Csv2Json))

	// Keep the program running
	select {}
}
