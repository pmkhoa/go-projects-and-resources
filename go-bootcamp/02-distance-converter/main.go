package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	miToKm = 1.60934
)

func main() {
	fromUnit := os.Args[1]
	toUnit := os.Args[2]

	switch {
	case strings.HasSuffix(fromUnit, "mi"):
		fmt.Println("FROM MILES")
		switch toUnit {
		case "km":
			fromUnitValue, _ := strconv.ParseFloat(fromUnit[:len(fromUnit)-2], 64)
			fmt.Println(fromUnitValue * miToKm)
		case "ft":
			fmt.Println("To FT")
		case "m":
			fmt.Println("To Metters")
		}
	}
}
