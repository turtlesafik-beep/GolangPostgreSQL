package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Create("out/newfile.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("1")
}
