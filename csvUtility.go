package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

type TimeLineData struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

func readCsv(fname string) []TimeLineData {
	csvFile, err := os.Open(fname)

	if err != nil {
		log.Fatalln("Cant read file: {}", fname)
	}

	r := csv.NewReader(csvFile)
	r.Comma = ','
	r.FieldsPerRecord = -1

	layout := "2006-01"

	var records []TimeLineData

	for {

		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		_, terr := time.Parse(layout, record[0])

		if terr != nil {
			continue
		}

		i, ierr := strconv.Atoi(record[1])

		if ierr != nil {
			continue
		}

		records = append(records, TimeLineData{record[0], i})
	}

	return records

}
