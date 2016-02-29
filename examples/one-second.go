package main

import (
	"fmt"
	"time"

	"github.com/ifo/sanic"
)

func main() {
	w := sanic.NewWorker(0, 1288834974657, 4, 14, 41, time.Millisecond)

	start := time.Now()

	ids := []int64{}
	for time.Since(start).Seconds() < 1 {
		ids = append(ids, w.NextID())
	}

	fmt.Println(len(ids))

	if hasDuplicates(ids) {
		fmt.Println("ids have duplicates")
	} else {
		fmt.Println("ids do not have duplicates")
	}

	fmt.Println("=== first ===")
	first := w.IDString(ids[0])
	fmt.Println(first)

	fmt.Println("=== last ===")
	last := w.IDString(ids[len(ids)-1])
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
