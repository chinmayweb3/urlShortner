package main

import (
	"fmt"
	"time"
)

func main() {

	t1 := time.Now()
	fmt.Println("starting time: ", t1.Hour())
	fmt.Println("starting time: ", t1.Minute())
	fmt.Println("starting time: ", t1.Second())

	time.Sleep(10 * time.Second)

	show := time.Since(t1)
	fmt.Println("time: ", show)
	fmt.Println("time: ", show.Hours())
	fmt.Println("time: ", show.Minutes())
	fmt.Println("time: ", show.Seconds())
}
