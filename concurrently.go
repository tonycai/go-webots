package main

import "fmt"
import "time"
import "math/rand"
import "runtime"

func main() {
	// Figure out how many CPUs are available and tell Go to use all of them
	numThreads := runtime.NumCPU()
	runtime.GOMAXPROCS(numThreads)

	// Number of random ints to generate
	var numIntsToGenerate = 100000000
	// Number of ints to be generated by each spawned goroutine thread
	var numIntsPerThread = numIntsToGenerate / numThreads

	// Channel for communicating from goroutines back to main function
	ch := make(chan []int)

	fmt.Printf("Initiating single-threaded random number generation.\n")
	startSingleRun := time.Now()
	// Generate all of the ints from a single goroutine, retrieve the expected
	// number of ints from the channel and put in target slice
	go makeRandomNumbers(numIntsToGenerate, ch)
	singleThreadIntSlice := <-ch
	elapsedSingleRun := time.Since(startSingleRun)
	fmt.Printf("Single-threaded run took %s\n", elapsedSingleRun)

	fmt.Printf("Initiating multi-threaded random number generation.\n")

	multiThreadIntSlice := make([][]int, numThreads)
	startMultiRun := time.Now()
	// Run the designated number of goroutines, each of which generates its
	// expected share of the total random ints, retrieve the expected number
	// of ints from the channel and put in target slice
	for i := 0; i < numThreads; i++ {
		go makeRandomNumbers(numIntsPerThread, ch)
	}
	for i := 0; i < numThreads; i++ {
		multiThreadIntSlice[i] = <-ch
	}
	elapsedMultiRun := time.Since(startMultiRun)
	fmt.Printf("Multi-threaded run took %s\n", elapsedMultiRun)
	//To avoid not used warning
	fmt.Print(len(singleThreadIntSlice))
}

func makeRandomNumbers(numInts int, ch chan []int) {
	source := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(source)
	result := make([]int, numInts)
	for i := 0; i < numInts; i++ {
		result[i] = generator.Intn(numInts * 100)
	}
	ch <- result
}
