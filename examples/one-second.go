package main

import (
	"fmt"
	"time"

	"github.com/ifo/sanic"
)

func main() {
	w := sanic.NewWorker10(0)

	ids := []int64{}
	for start := time.Now(); time.Since(start).Seconds() < 1; {
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
