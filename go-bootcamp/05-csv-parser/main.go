package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type State struct {
	id               int
	name             string
	abbreviation     string
	censusRegionName string
}

func parseState(columns map[string]int, record []string) (*State, error) {
	column := columns["id"]
	id, err := strconv.Atoi(record[column])
	if err != nil {
		log.Fatalln("Cannot convert string to int")
		return nil, err
	}
	name := record[columns["name"]]
	abbreviation := record[columns["abbreviation"]]
	censusRegionName := record[columns["censusRegionName"]]
	return &State{
		id:               id,
		name:             name,
		abbreviation:     abbreviation,
		censusRegionName: censusRegionName,
	}, nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Provide more args")
	}
	f, err := os.Open("state_table.csv")
	if err != nil {
		log.Fatalln("Cannot open file")
	}
	defer f.Close()

	stateLookup := map[string]*State{} // map of string and state for looking up state by abbr
	csvRead := csv.NewReader(f)
	columns := make(map[string]int) // using columns to hold field
	for rowCount := 0; ; rowCount++ {
		record, err := csvRead.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln("Error while reading")
		}

		// First row is the label
		if rowCount == 0 {
			for index, column := range record {
				columns[column] = index
			}
		} else {
			state, err := parseState(columns, record)
			if err != nil {
				log.Fatalln("Error while parsing state")
			}
			stateLookup[state.abbreviation] = state
		}

	}

	abbreviation := os.Args[1]
	state, ok := stateLookup[abbreviation]
	if !ok {
		log.Fatalln("Cannot find state")
	}
	fmt.Println(state)
}
