package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	miToKm = 1.60934
)

func main() {
	number, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("<!DOCTYPE html>")
	fmt.Println("<html>")
	fmt.Println("<head></head>")
	fmt.Println("<body")

	fmt.Printf("Miles: %f\n", number)
	fmt.Printf("Km: %f\n", number*miToKm)
	fmt.Println("</body>")
	fmt.Println("</html>")

}
