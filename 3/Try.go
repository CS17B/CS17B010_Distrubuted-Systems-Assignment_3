package main

import (
	"fmt"
	"sync"
)

var ledger int = 0
var wg sync.WaitGroup

func worker1(a int, c chan int) {
	c <- a
	ledger = <-c
	wg.Done()

}

func worker2(b int, c chan int) {
	c <- b
	ledger = <-c
	wg.Done()

}

func worker3(d int, c chan int) {
	c <- d
	ledger = <-c
	wg.Done()
}

func main() {
	c := make(chan int)
	wg.Add(3)
	go worker1(5, c)
	go worker2(3, c)
	go worker3(4, c)

	wg.Wait()
	fmt.Printf("%d\n", ledger)
}
