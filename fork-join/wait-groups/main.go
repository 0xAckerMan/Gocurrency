package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	now := time.Now()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		task1()
		task2()
		task3()
		task4()
	}()
	wg.Wait()
	fmt.Println("Elapsed, ", time.Since(now))
}

func task1() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Task 1")
}

func task2() {
	time.Sleep(400 * time.Millisecond)
	fmt.Println("Task 2")
}

func task3() {
	fmt.Println("Task 3")
}

func task4() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Task 4")
}
