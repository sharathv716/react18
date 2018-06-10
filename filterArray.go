package main

import "fmt"

var evenArray []int

func filterPositions(arr []int, index int){
  if index <= len(arr) - 1 && index%2 == 0{
    evenArray = append(evenArray, arr[index])
    filterPositions(arr, index+1)
  }else if index <= len(arr) - 1{
    filterPositions(arr, index+1)
  }
}

func main() {
  arr := []int{1, 2, 3, 4, 5, 6}
  filterPositions(arr, 0)
  fmt.Println(evenArray)
}
