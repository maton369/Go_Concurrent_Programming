package main

import (
	listing6_10 "Go_Concurrent_Programming/chapter6/listing6.10"
	"fmt"
	"math/rand"
	"time"
)

const matrixSize = 1200

func generateRandMatrix(matrix *[matrixSize][matrixSize]int) {
	for row := 0; row < matrixSize; row++ {
		for col := 0; col < matrixSize; col++ {
			matrix[row][col] = rand.Intn(10) - 5
		}
	}
}

func rowMultiply(matrixA, matrixB, result *[matrixSize][matrixSize]int,
	row int, barrier *listing6_10.Barrier) {
	for {
		barrier.Wait()
		for col := 0; col < matrixSize; col++ {
			sum := 0
			for i := 0; i < matrixSize; i++ {
				sum += matrixA[row][i] * matrixB[i][col]
			}
			result[row][col] = sum
		}
		barrier.Wait()
	}
}

func main() {
	var matrixA, matrixB, result [matrixSize][matrixSize]int
	barrier := listing6_10.NewBarrier(matrixSize + 1)
	for row := 0; row < matrixSize; row++ {
		go rowMultiply(&matrixA, &matrixB, &result, row, barrier)
	}
	start := time.Now()
	for i := 0; i < 4; i++ {
		generateRandMatrix(&matrixA)
		generateRandMatrix(&matrixB)
		barrier.Wait()
		barrier.Wait()
	}
	fmt.Println("Complete in", time.Since(start))
}