package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// Go routines and wait groups

func Progression(start int, multiplier int, wg *sync.WaitGroup) {
	maxLoops := 10
	val := start
	prefix := "Progression (" + strconv.Itoa(start) + "," + strconv.Itoa(multiplier) + ") "

	for i := 0; i < maxLoops; i++ {
		fmt.Printf(prefix + strconv.Itoa(i) + ": " + strconv.Itoa(val) + "\n")
		val *= multiplier
		time.Sleep(250 * time.Millisecond)
	}
	wg.Done()
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
	for {
		i, ok := <-ch
		if ok == false {
			fmt.Printf("channel closed")
			return
		}
		fmt.Printf("Pulled off %d\n", i)
		time.Sleep(1000 * time.Millisecond)
	}
}

func BufferChannels() {
	maxbuffer := 5
	chBuffered := make(chan int, maxbuffer)
	go PushChannel(chBuffered)
	PopChannel(chBuffered)
}
