package main

var evenArray []int

func filterPositions(arr []int, index int) {
	if index <= len(arr)-1 && index%2 == 0 {
		evenArray = append(evenArray, arr[index])
		filterPositions(arr, index+1)
	} else if index <= len(arr)-1 {
		filterPositions(arr, index+1)
	}
}
