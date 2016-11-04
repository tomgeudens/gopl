/*
description :
  fetchall fetches URLs in parallel and reports their times and sizes

author :
  Tom Geudens (https://github.com/tomgeudens/)

modified :
  2016/11/04
*/
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func fetch(url string, ch chan<- string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("fetchall: io.Copy %s: %v", url, err)
		return
	}

	timeElapsedSeconds := time.Since(start).Seconds()

	ch <- fmt.Sprintf("%.2fs  %7d  %s", timeElapsedSeconds, nbytes, url)
}

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}

	for range os.Args[1:] { // we just need the correct number of iterations here
		fmt.Println(<-ch) // receive from channel ch
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
