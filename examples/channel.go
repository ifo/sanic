package main

import (
	"fmt"
	"time"

	"github.com/ifo/sanic"
)

func main() {
	idChan := make(chan int64, 10000)
	w := sanic.NewWorker(0, 0, 1288834974657, 5, 5, 12, 41, idChan)
	start := time.Now()

	w.GenIDsForever()

	ids := []int64{}
	for time.Since(start).Seconds() < 1 {
		ids = append(ids, <-idChan)
	}

	fmt.Println(len(ids))

	if hasDuplicates(ids) {
		fmt.Println("ids have duplicates")
	} else {
		fmt.Println("ids do not have duplicates")
	}

	fmt.Println("=== first ===")
	first, _ := sanic.IntToString(ids[0])
	fmt.Println(first)

	fmt.Println("=== last ===")
	last, _ := sanic.IntToString(ids[len(ids)-1])
	fmt.Println(last)
}

func hasDuplicates(ids []int64) bool {
	set := map[int64]struct{}{}
	for i := range ids {
		if _, has := set[ids[i]]; has {
			return true
		}
		set[ids[i]] = struct{}{}
	}
	return false
}
