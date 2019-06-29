package main

import "sync"

/*func main() { // will write code concurrently

	var wg sync.WaitGroup
	m :=  make(map[string]int)

	for x := 0; x < 10000; x++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			m["hello"] = 1
		}(&wg)
	}
	wg.Wait()
}*/

type Counter struct {
	mx sync.RWMutex
	m  map[string]int
}

func main() {

	var wg sync.WaitGroup
	var counter = Counter{m: make(map[string]int)}

	for x := 0; x < 10000; x++ {
		wg.Add(1)
		go func(counter *Counter) {
			defer wg.Done()
			counter.mx.Lock()
			counter.m["hello"] = 1
			counter.mx.Unlock()
		}(&counter)
	}
	wg.Wait()

}
