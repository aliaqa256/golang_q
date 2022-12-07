package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	// "math/rand"
)

func main() {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)


	ch := make(chan int )
	fmt.Println("producer:started")

	go consumer1(ch)
	go consumer2(ch)
	for i := 0; i < 3; i++ {
		fmt.Println("producer:send ", i)
		go prodcer1(ch, i)
		fmt.Println("producer:send ", i, " done")

	}

	print("producer:done ")


	<-sigs
	print("producer:exit ")
}

func prodcer1(ch chan int, i int) {
	fmt.Println("producer:send ", i, " to channel")
	for {
		ch <- rand.Intn(100)
	}
}

func consumer1(ch chan int) {
	fmt.Println("consumer:started")
	for {
		select {
		case i := <-ch:
			fmt.Println(i, "---------------------------------------- ")
		default:
			continue
		}
	}
}

func consumer2(ch chan int) {
	fmt.Println("consumer:started")
	for {
		select {
		case i := <-ch:
			fmt.Println(i,"++++++++++++++++++++++++++++++++++++++++++++")
		default:
			continue
		}
	}
}

