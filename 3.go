package main

import (
	"fmt"
	"sync"
)

var ledger int = 0
var wg sync.WaitGroup

var ch chan int

func worker1() {
	ch <- 1
	ledger = 5
	<-ch
	wg.Done()
}

func worker2() {
	ch <- 1
	ledger = 4
	<-ch
	wg.Done()
}

func worker3() {
	ch <- 1
	ledger = 3
	<-ch
	wg.Done()
}

func main() {
	ch = make(chan int)
	wg.Add(3)
	go worker1()
	go worker2()
	go worker3()
	wg.Wait()
	fmt.Printf("%d\n", ledger)
}
