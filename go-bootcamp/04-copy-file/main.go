package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func my_cp(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		log.Fatalln("Cannot open file")
	}
	defer srcFile.Close()

	bufferSlice, err := ioutil.ReadAll(srcFile)
	if err != nil {
		return fmt.Errorf("Cannot parse src file")
	}

	dstFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("Cannot create dst file")
	}
	defer srcFile.Close()

	_, err = dstFile.Write(bufferSlice)
	if err != nil {
		return fmt.Errorf("Cannot write to file")
	}

	return nil
}

func cp(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		log.Fatalln("Cannot open file")
	}
	defer srcFile.Close()

	if err != nil {
		return fmt.Errorf("Cannot parse src file")
	}

	dstFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("Cannot create dst file")
	}
	defer srcFile.Close()

	_, err = io.Copy(dstFile, srcFile)

	if err != nil {
		return fmt.Errorf("Cannot write to file")
	}
	return nil
}

func main() {
	if len(os.Args) < 3 {
		log.Fatalln("Copy need to have at least src & dst")
	}

	src := os.Args[1]
	dst := os.Args[2]

	err := cp(src, dst)
	if err != nil {
		log.Fatalln("Cannot copy")
	}
	fmt.Println("File copied successful")
}
