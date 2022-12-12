package main

import (
	"fmt"
	"net/http"
	"sync"
)

var urls []string = []string{"https://google.com", "https://mongodb.com", "https://github.com", "https://heroku.com", "https://linkedin.com"}

func main() {

	sem := make(chan int, 4)
	errChan := make(chan string, 4)
	var wg sync.WaitGroup
	wg.Add(len(urls))
	for _, url := range urls {
		go fetch(sem, url,  &wg, errChan)
	}
	wg.Wait()
	close(errChan)

}

func fetch(sem chan int, url string, wg *sync.WaitGroup, errChan chan string) {
	sem <- 1
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		errChan <- err.Error()
	}
 
	fmt.Println("===========================================================================================================")
	fmt.Println(resp)
	fmt.Println("===========================================================================================================")
	<-sem
}
