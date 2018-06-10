package main

import "fmt"

var (
  count1 int
  count2 int
)

func sequenceOfColors(str string, index int) bool{
  if string(str[index]) == "R" || string(str[index]) == "G" && index < len(str)-1{
    count1++
    index++
    sequenceOfColors(str[index:], index)
  }else if string(str[index]) == "Y" || string(str[index]) == "B" && index < len(str)-1{
    count2++
    index++
    sequenceOfColors(str[index:], index)
  }else if index < len(str) - 1{
    index++
    sequenceOfColors(str[index:], index)
  }
  if (count1+count2)%2 == 0{
    return true
  }else{
    return false
  }

}

func main(){
  fmt.Println(sequenceOfColors("RGGRYB", 0))
}
