package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	csvFileName = "data/example.csv"
	outFileName = "data/output.txt"
)

type Data struct {
	key   string
	value string
}

func main() {

	csvFile, err := os.Open(csvFileName)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, col := range csvLines {
		d := Data{
			key:   col[0],
			value: col[1],
		}

		writeOutput(d)
	}

}

func writeOutput(d Data) {
	f, err := os.OpenFile(outFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	data := trimData(d)

	// remove unwanted characters from value
	for _, r := range rules {
		data.value = strings.Replace(data.value, r.character, r.replacement, -1)
	}

	f.WriteString("\"" + data.key + "\": " + data.value + ", \n")
}

func trimData(d Data) *Data {
	key := strings.TrimSpace(d.key)
	value := strings.TrimSpace(d.value)

	return &Data{
		key,
		value,
	}
}
