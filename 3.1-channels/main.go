package main

import (
	"fmt"
)

func generateNumbers(ch chan int) {
	for i := 2; ; i++ {
		fmt.Println("generateNumbers", i)
		ch <- i
	}
}

func filterPrimes(in, out chan int, prime int, done chan struct{}) {
	for {
		select {
		case num := <-in:
			fmt.Println("filterPrimes", num)
			if num%prime != 0 {
				out <- num
			}
		case <-done:
			close(out)
			return
		}
	}
}

func main() {
	const n = 10
	ch := make(chan int)
	done := make(chan struct{})

	go generateNumbers(ch)

	for i := 0; i < n; i++ {
		prime := <-ch
		fmt.Println(prime)

		ch1 := make(chan int)
		go filterPrimes(ch, ch1, prime, done)
		ch = ch1
	}
	close(done)
}
