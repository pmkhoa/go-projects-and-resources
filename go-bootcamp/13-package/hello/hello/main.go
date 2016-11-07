package main

import (
	"fmt"

	"github.com/pmkhoa/go-bootcamp/13-package/hello"
	"github.com/ttacon/chalk"
)

func main() {
        hello.Hello()
	fmt.Println(chalk.Red, "use chalk package", chalk.ResetColor)

}
