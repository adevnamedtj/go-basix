package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println("boot test")
	c := make(chan string)
	for i := 0; i <= 5; i++ {

		go run(c, i)
	}
	for i := 0; i <= 5; i++ {
		fmt.Println(<-c)
	}
}

func run(c chan string, x int) {
	timeStart := time.Now()

	for i := 0; i <= 100; i++ {
		fmt.Println(x, i)
	}
	diff := time.Now().Sub(timeStart)
	c <- ("done " + strconv.Itoa(x) + " in " + diff.String())
}
