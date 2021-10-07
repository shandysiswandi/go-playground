package main

import "sync"

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 76; i++ {
		wg.Add(1)
		go processExcelColumn(&wg)
	}

	wg.Wait()
}

func processExcelColumn(wg *sync.WaitGroup) {
	defer wg.Done()
	// logic
}
