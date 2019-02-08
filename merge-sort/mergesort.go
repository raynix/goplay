package main

import (
  "strconv"
  "os"
  "log"
  "fmt"
  "math/rand"
)

func MergeSort(a []int) []int {
  length := len(a)
  if length <= 1 {
    return a
  }
  middle := length / 2
  left := MergeSort(a[:middle])
  right := MergeSort(a[middle:])
  return Merge(left, right)
}

func Merge(left []int, right []int) (result []int){
  result = []int {}
  for ;len(left) > 0 || len(right) > 0; {
    switch{
    case len(left) == 0:
      result = append(result, right[0])
      right = right[1:]
    case len(right) == 0:
      result = append(result, left[0])
      left = left[1:]
    default:
      if left[0] > right[0] {
        result = append(result, right[0])
        right = right[1:]
      }else{
        result = append(result, left[0])
        left = left[1:]
      }
    }
  }
  return
}

func main() {
  length := 100
	if len(os.Args) > 1 {
		l, err := strconv.Atoi(os.Args[1])
    if err != nil {
  		log.Fatal(err)
    }
    length = l
	}
  a := []int{}
  for i := 0; i< length; i++ {
      a = append(a, rand.Intn(length))
  }
  fmt.Println(a)
  fmt.Println(MergeSort(a[:]))
}
