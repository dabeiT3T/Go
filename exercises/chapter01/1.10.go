package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {

	start := time.Now()
	url := os.Args[1]
	ch := make(chan string)

	for i := 0; i < 2; i++ {
		go fetch(url, ch, i)
		// block here, so that the server can cache the response
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string, i int) {

	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	filename := "body_" + strconv.Itoa(i)
	f, err := os.Create(filename)
	if err != nil {
		ch <- fmt.Sprintf("While creating %s: %v", filename, err)
		return
	}

	nbytes, err := io.Copy(f, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("While reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
