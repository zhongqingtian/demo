package sync

import (
	"log"
	"sync"
)

func WgGro() bool {
	wg1 := &sync.WaitGroup{}
	wg2 := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg1.Add(1)
		go worker(wg1, i)
	}
	wg1.Wait()

	for i := 0; i < 5; i++ {
		wg1.Add(1)
		go worker(wg1, i)
	}
	wg1.Wait()
	return wg1 == wg2 // Prints true
}

func worker(wg *sync.WaitGroup, i int) {
	defer wg.Done()
	log.Print(wg)
	log.Print("i=", i)
}
