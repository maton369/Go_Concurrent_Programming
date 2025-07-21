package main

import (
	listing4_12 "Go_Concurrent_Programming/chapter4/listing4.12"
	"fmt"
	"time"
)

func main() {
    rwMutex := listing4_12.ReadWriteMutex{}
    for i := 0; i < 2; i++ {
        go func() {
            for {
                rwMutex.ReadLock()
                time.Sleep(1 * time.Second)
                fmt.Println("Read done")
                rwMutex.ReadUnlock()
            }
        }()
    }
    time.Sleep(1 * time.Second)
    rwMutex.WriteLock()
    fmt.Println("Write finished")
}