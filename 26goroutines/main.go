package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup // pointer
var mut sync.Mutex    // pointer

func main() {
	// go greeter("hello")
	// greeter("world")

	websitelist := []string{
		"https://google.co.in",
		"https://google.com",
		"https://amazon.in",
		"https://github.com",
	}
	for _, web := range websitelist {
		go getStatusCode(web)

		// Keep on adding go co routines to a WaitGroup
		wg.Add(1)
	}

	// Waits for the go routines to complete
	wg.Wait()

	// Print the Signals
	fmt.Println(signals)

}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)

// 	}

// }

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}

}
