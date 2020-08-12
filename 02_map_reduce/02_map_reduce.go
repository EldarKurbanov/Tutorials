package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	data := getSlice(100_000_000, 100)

	start := time.Now()

	result := make(chan int)

	// todo можно попробовать менять это значение и таким образов создавать больше или меньше горутин и посмотреть с какой скоростью выполняется код
	//const chunkSize = 1_000_000_000_000_000_000
	const chunkSize = 1000000

	// pattern блокирующий канал результатов
	go processSlice(data, chunkSize, result)

	sum := <-result
	fmt.Println("The result for concurrent case is", sum, time.Since(start))

	// direct sum
	start = time.Now()
	sum = 0
	for _, v := range data {
		sum += v
	}
	fmt.Println("The result sequential case is", sum, time.Since(start))
}

func processSlice(input []int, chunkSize int, res chan<- int) {
	sliceLen := len(input)

	chunks := sliceLen / chunkSize
	if len(input)%chunkSize != 0 {
		chunks++
	}

	results := make(chan int, chunks)
	var current, next int

	for i := 0; i < chunks; i++ {
		current, next = i*chunkSize, (i+1)*chunkSize
		if i == chunks-1 {
			next = sliceLen
		}

		// pattern workers по чанкам(кусочкам) данных
		go processPart(input, current, next, results)
	}

	result := 0
	i := 0
	// pattern канал множества результатов из множества горутин
	// pattern map-reduce
	for chunkResult := range results {
		result += chunkResult

		i++
		if i == chunks {
			// завершение работы канала для выхода из for
			close(results)
		}
	}

	res <- result
}

func processPart(input []int, from, to int, res chan<- int) {
	sum := 0
	for i := from; i < to; i++ {
		sum += input[i]
	}
	res <- sum
}

func getSlice(n int, max int32) []int {
	ns := make([]int, n)
	for i := range ns {
		ns[i] = int(rand.Int31n(max))
	}

	return ns
}
