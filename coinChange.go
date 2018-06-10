package main

import (
	"fmt"
)


var chosen []int

var cache map[int][]int


func coinChange(coins []int, target int, arr []int){

  if len(cache[target]) == 0 {
    cache[target] = arr
  }

  if target == 0{
    fmt.Println("the current set >>> ", arr)
  }

  for _, v := range coins{
    if target - v >= 0{
      arr = append(arr, v)
      coinChange(coins, target-v, arr)
      arr = arr[:len(arr) - 1]
    }
  }
}

func main() {
  cache = make(map[int][]int)
	coins := []int{25, 10, 5, 1}
	target := 100
	coinChange(coins, target, chosen)
  fmt.Println(cache)

}
