package main

import (
  "fmt"
)
var i, j, count int
var mystring []string

func max(s1, s2 int) int{
  if s1 < s2{
    return s2
  }
  return s1
}

func lcs(str1, str2 string) int{
  count++
  fmt.Println("new change1")
  fmt.Println("new change2")
  if i > len(str1) - 1  || j > len(str2) - 1{
    return 0
  }
  if str1[i] == str2[j]{
    return 1 + lcs(str1[i+1:], str2[j+1:])
  }else{
    return max(lcs(str1[i+1:], str2[j:]), lcs(str1[i:], str2[j+1:]))
  }
}

func main(){
  fmt.Println(lcs("AGGTAB","GXTXAYB"))
  fmt.Println("Number of calls ", count)
  fmt.Println(mystring)
  fmt.Println("On Testing Branch")
}
