package main

import "fmt"

// START OMIT
func worker(id int, jobs <-chan int, results chan<- int) {
	select {
	case j := <-jobs:
		results <- j * 2
	default:
		fmt.Println("no tengo nada para hacer")
	}
}

func main() {
	jobs, results := make(chan int, 100), make(chan int, 100)
	for w := 1; w <= 5; w++ {
		go worker(w, jobs, results)
	}

	values := []int{1, 5, 7}
	for _, v := range values {
		jobs <- v
	}
	// close(jobs) // no cierro el channel adrede
	// END OMIT

	for _ = range values {
		fmt.Println("result:", <-results)
	}
}
