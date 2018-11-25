package main

import (
    "fmt"
    "strings"
    "bufio"
    "os"
    "sync"
    "sync/atomic"
)

func formatOutput(url string, count int) string{
    return fmt.Sprintf("%s:%d", url, count)
}

func lookupWorker(word string, input <-chan string, total *int64){
    for url := range input {
        // strings.Count doesnt count overlaps
        urlCount := strings.Count(url, word)
        fmt.Println(formatOutput(url, urlCount))
        atomic.AddInt64(total, int64(urlCount))
    }
}


func startWorker(wg *sync.WaitGroup, word string, input <-chan string, total *int64){
    // didnt want to mix wg logic with lookup
    lookupWorker(word, input, total)
    wg.Done()
}


func CountWord(scanner *bufio.Scanner, word string, workersLimit int, inputBuffer int){
    wg := &sync.WaitGroup{}

    // it is possible to count goroutines by runtime.NumGoroutine()
    // but i like currentWorkers way
    currentWorkers := 0

    // channel for workers input
    input := make(chan string, inputBuffer)
    var total int64

    for scanner.Scan() {
        url := scanner.Text()
        if (currentWorkers < workersLimit){
            wg.Add(1)
            // maybe closure?
            go startWorker(wg, word, input, &total)
            currentWorkers += 1
        }
        input <- url
    }

    if scanner.Err() != nil {
        fmt.Println("Reading goes wrong")
    }

    close(input)
    wg.Wait()
    fmt.Println("Total:", total)
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    // word for search
    word := "golang"
    // in total will be workersLimit+1 goroutines (+1 for main)
    workersLimit := 5
    // size of input chan buffer
    inputBuffer := 10

    CountWord(scanner, word, workersLimit, inputBuffer)
}


// echo -e 'https://golang.org\nhttps://url.com\nhttps://example.org\nhttps://golanggolang.org\nhttps://golanggolang.org\nhttps://golanggolang.org\nhttps://golanggolang.org\nhttps://golanggolang.org' | go run urlFinder.go