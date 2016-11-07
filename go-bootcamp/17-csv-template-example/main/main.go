package main

import (
	"encoding/csv"
	"net/http"
	"os"
        "html/template"
	"github.com/pmkhoa/go-bootcamp/17-csv-template-example/record"
)

func renderCSV(res http.ResponseWriter, req *http.Request) {
	f, err := os.Open("table.csv")
	if err != nil {
		http.Error(res, err.Error(), 500)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	rows, err := csvReader.ReadAll()
	if err != nil {
		http.Error(res, err.Error(), 500)
	}

	records := make([]record.Record, 0, len(rows))
	for index, row := range rows {
		if index == 0 {
			continue
		} else {
			record := record.MakeRecord(row)
			records = append(records, record)
		}
	}

        tpl, err := template.ParseFiles("index.gotpl")
        if err != nil {
		http.Error(res, err.Error(), 500)
        }
        err = tpl.Execute(res, records)
        if err != nil {
		http.Error(res, err.Error(), 500)
        }
}

func main() {
	http.HandleFunc("/", renderCSV)
	http.ListenAndServe(":9000", nil)

}
