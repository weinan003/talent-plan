package main

import "sync"

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.


func Conquer(src []int64, anchor int) {
	lhsid := 0
	rhsid := anchor

	for lhsid < anchor && rhsid < len(src) {
		for lhsid < anchor && src[lhsid] z
	}
}

func Divide(src []int64,parentWG *sync.WaitGroup) {

	if(len(src) >= 2) {
		var wg sync.WaitGroup

		wg.Add(2)
		go Divide(src[:len(src)/2], &wg)
		go Divide(src[len(src)/2:], &wg)
		wg.Wait()

		Conquer(src, len(src) / 2)
	}

	parentWG.Done()
}

func MergeSort(src []int64) {

	if len(src) < 2 {
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)
	Divide(src, &wg)
	wg.Done()
}
