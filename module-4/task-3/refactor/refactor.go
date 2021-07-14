package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func readFile(path string) []byte {
	fd, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}

	raw, err := ioutil.ReadAll(fd)
	if err != nil {
		panic(err)
	}

	return raw
}

func countLines(raw []byte) int {
	count := 0
	for i := 0; i < len(raw); i++ {
		if raw[i] == '\n' {
			count++
		}
	}

	return count
}

func main() {
	path := "./README.md"

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	raw := readFile(path)

	fmt.Println(countLines(raw))
}
