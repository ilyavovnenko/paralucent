package main

import (
	"fmt"
	"sync"
)

func deposit(b *int, n int, wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()

	*b += n

	wg.Done()
}

func withdraw(b *int, n int, wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()

	*b -= n

	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	balance := 100

	for i := 0; i < 100; i++ {
		wg.Add(2)
		go deposit(&balance, i, &wg, &mu)
		go withdraw(&balance, i, &wg, &mu)
	}
	wg.Wait()

	fmt.Println("Final balance value:", balance)
}
