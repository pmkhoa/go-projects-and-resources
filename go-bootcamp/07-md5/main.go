package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func main() {
	hash := md5.New()
	// This is used to hash a file
	// f, err := os.Open(os.Args[1])
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()
	// io.Copy(hash, f)

	// Hash a string input from user
	io.WriteString(hash, os.Args[1])

	fmt.Printf("%x\n", hash.Sum(nil))
}
