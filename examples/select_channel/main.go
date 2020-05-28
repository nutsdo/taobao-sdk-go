package main

import (
	"fmt"
	"time"
)

func main() {
	reqC := make(chan bool, 1)
	reqB := make(chan bool, 1)
	go func() {
		time.Sleep(time.Second*10)
		reqC <- true
	}()

	go func() {
		time.Sleep(time.Second*15)
		reqB <- true
	}()
	fmt.Println("等待接受数据...")
	select {
	case <-reqC:
		fmt.Println("131232")
	case <-reqB:
		fmt.Println("bbbbbbb")

	}
}