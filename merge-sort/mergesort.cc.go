package main

import (
  "strconv"
  "os"
  "log"
  "fmt"
  "math/rand"
)
var sem = make(chan struct{}, 4)
func MergeSort(a []int, result chan []int) {
  length := len(a)
  if length <= 1 {
    result <- a
    return
  }
  middle := length / 2
  split_channel := make(chan []int, 2)
  select{
  case sem <- struct{}{}:
    go func(){
      MergeSort(a[:middle], split_channel)
      MergeSort(a[middle:], split_channel)
      <- sem
    }()
  default:
    MergeSort(a[:middle], split_channel)
    MergeSort(a[middle:], split_channel)
  }

  result <- Merge(<-split_channel, <-split_channel)
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
  top := make(chan []int, 1)
  MergeSort(a, top)
  fmt.Println(<- top)
}
