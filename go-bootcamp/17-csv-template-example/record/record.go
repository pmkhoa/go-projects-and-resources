package record

import "strconv"

type Record struct {
	Date string
	Open float64
}

func MakeRecord(row []string) Record {
	open, _ := strconv.ParseFloat(row[1], 64)
	return Record{
		Date: row[0],
		Open: open,
	}
}
