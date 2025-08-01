package main

import (
	listing6_7 "Go_Concurrent_Programming/chapter6/listing6.7"
	"fmt"
)

func doWork(id int, wg *listing6_7.WaitGrp) {
    fmt.Println(id, "Done working ")
    wg.Done()
}

func main() {
    wg := listing6_7.NewWaitGrp()
    for i := 1; i <= 4; i++ {
        wg.Add(2)
        go doWork(i, wg)
        go doWork(i, wg)
    }
    wg.Wait()
    fmt.Println("All complete")
}