package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	end := time.Now()
	fmt.Println(end.Sub(start))

	start = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	end = time.Now()
	fmt.Println(end.Sub(start))
}

// $ go run 1.3.go param0 param1 param2 param3
// > param0 param1 param2 param3
// > 19.794µs
// > param0 param1 param2 param3
// > 2.288µs
