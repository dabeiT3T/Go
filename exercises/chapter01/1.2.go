package main

import (
	"fmt"
	"os"
)

func main() {
	for idx, elem := range os.Args[1:] {
		fmt.Println(idx, elem)
	}
}
