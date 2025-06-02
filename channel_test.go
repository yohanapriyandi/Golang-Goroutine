package belajar_golang_goroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Yohan Apriyandi"
		fmt.Println("Selesai mengirim data ke channel")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Yohan Apriyandi"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func OnLyIn(channel chan<- string) {
	time.Sleep(3 * time.Second)
	channel <- "Yohan Apriyandi"
}

func OnLyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}
func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnLyIn(channel)
	go OnLyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {

	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Yohan"
		channel <- "Fathar"
		channel <- "Taqy"
	}()

	go func() {
		fmt.Println("=============Print Channel=============")
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println("============End Print Channel==============")
	}()
	time.Sleep(5 * time.Second)
	fmt.Println("Selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 100; i++ {
			channel <- "Perualangan ke " + strconv.Itoa(i)
		}
		defer close(channel)
	}()

	for data := range channel {
		fmt.Println("Range ", data)
	}
	fmt.Println("Selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Ambil data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Ambil data dari channel 2", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}
}
