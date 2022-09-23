package main

import (
	"fmt"
	"sync"
)

func deposit(b *int, n int, mM *sync.Mutex, wg *sync.WaitGroup) {
	mM.Lock()
	*b += n
	mM.Unlock()
	wg.Done()
}

func withdraw(b *int, n int, mM *sync.Mutex, wg *sync.WaitGroup) {
	mM.Lock()
	*b -= n
	mM.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var mM sync.Mutex

	wg.Add(200)

	balance := 100

	for i := 0; i < 100; i++ {
		go deposit(&balance, i, &mM, &wg)
		go withdraw(&balance, i, &mM, &wg)
	}
	wg.Wait()

	fmt.Println("Final balance value:", balance)
}
