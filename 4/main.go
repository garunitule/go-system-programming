package main

import (
	"context"
	"fmt"
	"time"
)

func sub() {
	fmt.Println("sub is running")
	time.Sleep(2 * time.Second)
	fmt.Println("sub is finished")
}

func main() {
	fmt.Println("start sub()")
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		fmt.Println("sub() is finished")
		cancel()
	}()
	<-ctx.Done()
	fmt.Println("all tasks are finished")
}
