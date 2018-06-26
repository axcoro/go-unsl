package main

import "fmt"

// START OMIT
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "data:", j)
		results <- j * 2
	}
}

func main() {
	jobs, results := make(chan int, 100), make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	values := []int{1, 5, 7}
	for _, v := range values {
		jobs <- v
	}
	close(jobs)

	for _ = range values {
		fmt.Println("result:", <-results)
	}
}

// END OMIT
