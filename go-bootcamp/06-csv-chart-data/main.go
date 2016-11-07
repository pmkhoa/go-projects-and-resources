package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type record struct {
	Date string
	Open float64
}

func makeRecord(row []string) record {
	open, _ := strconv.ParseFloat(row[1], 64)
	return record{
		Date: row[0],
		Open: open,
	}
}

func main() {
	f, err := os.Open("table.csv")
	if err != nil {
		log.Fatalln("Cannot open file")
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalln("Cannot read file")
	}

	output := `<!DOCTYPE html>
                <htm>
                <head></head>
                <body>
                  <table>
                        <thead>
                                <tr>
                                <th>Date</th>
                                <th>Open</th>
                                </tr>
                        </thead>
                        <tbody>`

	for index, row := range records {
		if index == 0 {
			continue
		} else {
			record := makeRecord(row)
			output += `<tr>`
			output += `<td>` + record.Date + `</td>`
			output += `<td>` + fmt.Sprintf("%.2f", record.Open) + `</td>`
			output += `/<tr>`
		}
	}
	output += `</tbody></table> </body></html>`

	destFile, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("Cannot create file")
	}
	defer destFile.Close()

	_, err = destFile.WriteString(output)
	if err != nil {
		log.Fatalln("Cannot write data to file")
	}
}
