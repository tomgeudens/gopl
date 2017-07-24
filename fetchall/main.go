/*
description :
  fetchall fetches URLs in parallel and reports duration, size and status

author :
  Tom Geudens (https://github.com/tomgeudens/)

modified :
  2017/07/24
*/
package main

import "fmt"
import "io"
import "io/ioutil"
import "net/http"
import "os"
import "time"

func fetch(url string, ch chan<- string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		// send error to the channel
		ch <- fmt.Sprintf("fetchall: fetch: %v", url, err)
		return
	}

	written, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		// send error to the channel
		ch <- fmt.Sprintf("fetchall: fetch: reading %s: %v", url, err)
		return
	}

	elapsedSeconds := time.Since(start).Seconds()

	// send result to the channel
	ch <- fmt.Sprintf("%.2fs  %7d  %s (%d)", elapsedSeconds, written, url, resp.StatusCode)
}

func main() {
	start := time.Now()

	ch := make(chan string)

	// loop through the arguments (which should be urls)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}

	// get the results (same number as we had arguments)
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel
	}

	elapsedSeconds := time.Since(start).Seconds()

	// print total elapsed time
	fmt.Printf("%.2fs total elapsed time\n", elapsedSeconds)
}
