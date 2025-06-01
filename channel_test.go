package belajar_golang_goroutine

import (
	"fmt"
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

func nlTestChannelAsParameter(t *testing.T) {
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
