package main

import (
	listing5_16 "Go_Concurrent_Programming/chapter5/listing5.16"
	"fmt"
)

func main() {
    semaphore := listing5_16.NewSemaphore(0)
    for i := 0; i < 50000; i++ {
        go doWork(semaphore)
        fmt.Println("Waiting for child goroutine")
        semaphore.Acquire()
        fmt.Println("Child goroutine finished")
    }
}

func doWork(semaphore *listing5_16.Semaphore) {
    fmt.Println("Work started")
    fmt.Println("Work finished")
    semaphore.Release()
}