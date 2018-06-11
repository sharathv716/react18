package main

import (
	"fmt"
)

var chosen []int

var cache map[int][]int

func coinChange(coins []int, target int, arr []int) {

	if len(cache[target]) == 0 {
		cache[target] = arr
	}

	if target == 0 {
		fmt.Println("the current set >>> ", arr)
	}

	for _, v := range coins {
		if target-v >= 0 {
			arr = append(arr, v)
			coinChange(coins, target-v, arr)
			arr = arr[:len(arr)-1]
		}
	}
}
