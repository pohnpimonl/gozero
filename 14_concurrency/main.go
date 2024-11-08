package main

import (
	"fmt"
	"sync"
	"time"
)

type job struct {
	name        string
	processTime time.Duration
}

func main() {

	/*
		https://go.dev/talks/2012/waza.slide#1

		Concurrency vs 	Parallelism
		Concurrency is about dealing with lots of things at once.

		Parallelism is about doing lots of things at once.

		Not the same, but related.

		Concurrency is about structure, parallelism is about execution.

		Concurrency provides a way to structure a solution to solve a problem that may (but not necessarily) be parallelizable.
			concurrent or parallel which one is better sometime it's depend on type of work.

			eg. if some work proces have to wait some other work process maybe use parallel not help.
	*/

	// mutex  use when you want to prevent mutual section - ( lock behavior, use with caution for performance issue)
	// channel use when you want to communicate between goroutine - ( use with caution, IMHO: complex to read and  debug for beginner).

	/* -- check number of cpu --*/
	// fmt.Printf("%d\n", runtime.NumCPU())

	/* sync wait group */

	/* with out go routine */
	// for _, v := range []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"} {
	// 	fmt.Printf("val is: %s\n", v)
	// 	time.Sleep(time.Millisecond * 500)
	// }

	/* go routine without wait group */
	// for _, v := range []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"} {
	// 	go func() {
	// 		fmt.Printf("val is: %s\n", v)
	// 		time.Sleep(time.Millisecond * 500)
	// 	}()

	// }

	/* goroutine with sync wait group */
	var wg sync.WaitGroup

	wg.Add(10)
	for _, v := range []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"} {
		go func() {
			fmt.Printf("val is: %s\n", v)
			time.Sleep(time.Millisecond * 500)
			wg.Done()
		}()
	}
	wg.Wait()

	/*--- worker pool ---*/
	// const numJobs = 10
	// const workerLimit = 3
	// jobs := make(chan job, numJobs)
	// results := make(chan string, numJobs)

	// for w := 1; w <= workerLimit; w++ {
	// 	go worker(w, jobs, results)
	// }

	// for j := 1; j <= numJobs; j++ {
	// 	jobs <- job{
	// 		name:        fmt.Sprintf("job %d", j),
	// 		processTime: time.Duration(time.Second * time.Duration(j)),
	// 	}
	// 	// if j == 5 {
	// 	// 	close(jobs) // panic send close channel
	// 	// }
	// }
	// close(jobs)

	// //for a := 1; a <= 11; a++ { // dead lock
	// for a := 1; a <= numJobs; a++ {
	// 	<-results
	// }

	/* -- select ---*/
	// 	c1 := make(chan int)
	// 	c2 := make(chan int)
	// 	done1 := make(chan bool)
	// 	done2 := make(chan bool)

	// 	go func() {
	// 		for i := 0; i < 5; i++ {
	// 			time.Sleep(time.Second * 1)
	// 			c1 <- i
	// 		}

	// 		done1 <- true

	// 	}()
	// 	go func() {
	// 		for i := 0; i < 5; i++ {
	// 			time.Sleep(time.Second * 2)
	// 			c2 <- i
	// 		}

	// 		done2 <- true

	// 	}()

	// 	var v1 bool
	// 	var v2 bool

	// outerloop:
	// 	for {
	// 		select {
	// 		case msg1 := <-c1:
	// 			fmt.Println("received c1", msg1)
	// 		case msg2 := <-c2:
	// 			fmt.Println("received c2", msg2)
	// 		case v1 = <-done1:
	// 		case v2 = <-done2: // not good practice
	// 			if v1 && v2 {
	// 				fmt.Printf("all finish\n")
	// 				break outerloop
	// 			}
	// 		}
	// 	}

}

func worker(id int, jobs <-chan job, results chan<- string) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j.name)
		time.Sleep(j.processTime)
		fmt.Println("worker", id, "finished job", j.name)
		results <- fmt.Sprintf("job %s run success", j.name)
	}
}
