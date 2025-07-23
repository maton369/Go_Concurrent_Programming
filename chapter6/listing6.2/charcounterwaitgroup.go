package main

import (
	"Go_Concurrent_Programming/chapter4/listing4_5"
	"fmt"
	"sync"
)

func main() {
    wg := sync.WaitGroup{}
    wg.Add(31)
    mutex := sync.Mutex{}
    var frequency = make([]int, 26)
    for i := 1000; i <= 1030; i++ {
        url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
        go func() {
            listing4_5.CountLetters(url, frequency, &mutex)
            wg.Done()
        }()
    }
    wg.Wait()
    mutex.Lock()
    for i, c := range listing4_5.AllLetters {
        fmt.Printf("%c-%d ", c, frequency[i])
    }
    mutex.Unlock()
}