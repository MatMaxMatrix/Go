package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var m = sync.RWMutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5", "id6", "id7", "id8", "id9", "id10"}
var result = []string{}

func main() {
	t0 := time.Now();
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\n Total execution time: %s", time.Since(t0))
	fmt.Printf("\n The Result are: %v", result)
}

func dbCall(i int) {
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	save(dbData[i])
	log()
	wg.Done()

}

func save(data string) {
	m.Lock()
	result = append(result, data)
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Printf("\n result: %v", result)
	m.RUnlock()
}
