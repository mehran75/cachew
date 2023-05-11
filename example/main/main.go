package main

import (
	"fmt"
	"github.com/mehran75/cachew/memorycache"
	"strconv"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(10)
	data := map[string]int{}

	for i := 0; i < 50; i++ {
		data[strconv.Itoa(i)] = i
	}

	cachew := memorycache.NewCache[string, int]()

	for k, v := range data {
		go cachew.Set(k, v, 10*time.Second)
	}

	time.Sleep(1 * time.Second)

	for k, v := range data {
		go func(k string, v int) {
			d, err := cachew.Get(k)
			if err != nil {
				panic(err)
			}
			fmt.Printf("cached val: %v --- real val: %v\n", *d, v)
			wg.Done()
		}(k, v)
	}

	wg.Wait()
}
