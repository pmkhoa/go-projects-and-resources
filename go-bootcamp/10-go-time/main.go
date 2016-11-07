package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	timeAsTime, _ := time.Parse("2006-01-02", "1988-02-13")
	fmt.Println(timeAsTime.Format("01/02/2006"))

	// Timediff
	from, to := os.Args[1], os.Args[2]

	fromTime, _ := time.Parse("2006-01-02", from)
	toTime, _ := time.Parse("2006-01-02", to)

	dur := toTime.Sub(fromTime)
	fmt.Println(int(dur / (time.Hour * 24)))

}
