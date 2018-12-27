package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// Go routines and wait groups

func Progression(start int, multiplier int, wg *sync.WaitGroup) {
	defer wg.Done()

	maxLoops := 10
	val := start
	prefix := "Progression (" + strconv.Itoa(start) + "," + strconv.Itoa(multiplier) + ") "

	for i := 0; i < maxLoops; i++ {
		fmt.Printf(prefix + strconv.Itoa(i) + ": " + strconv.Itoa(val) + "\n")
		val *= multiplier
		time.Sleep(250 * time.Millisecond)
	}
}

func PlayWithGoRoutines() {
	var wg sync.WaitGroup
	wg.Add(2)
	go Progression(1, 2, &wg)
	go Progression(2, 3, &wg)
	wg.Wait()
}

// Channels

func GenNumbers(chresult chan string) {
	stres := ""
	for i := 0; i < 10; i++ {
		stres += (strconv.Itoa(i) + ",")
		fmt.Printf("String so far: " + stres + "\n")
	}
	chresult <- stres
}

func PlayWithChannels() {
	chresult := make(chan string)
	go GenNumbers(chresult)

	fmt.Printf("Waiting on channel response\n")
	stres := <-chresult
	fmt.Printf("\nBack from channel: " + stres + "\n")
}

func _PlayWithChanels2(ch chan int) {
	//for i := range ch {
	//	fmt.Printf("%d\n", i)
	//}
	ch <- 5
}

func PlayWithChannels2() {
	ch := make(chan int)
	go _PlayWithChanels2(ch)
	//ch <- 5
	i := <-ch
	close(ch)
	i = <-ch
	fmt.Printf("%d", i)
}

//
// Buffered channels
//

func PushChannel(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Printf("Just Pushed %d\n", i)
		//time.Sleep(1000 * time.Millisecond)
	}
	close(ch)
}

func PopChannel(ch chan int) {
	/*
		for {
			i, ok := <-ch
			if ok == false {
				fmt.Printf("channel closed")
				return
			}
	*/
	// better way:
	for i := range ch {
		fmt.Printf("Pulled off %d\n", i)
		time.Sleep(1000 * time.Millisecond)
	}
	fmt.Println("exiting")
}

func BufferChannels() {
	maxbuffer := 5
	chBuffered := make(chan int, maxbuffer)
	go PushChannel(chBuffered)
	PopChannel(chBuffered)
}
