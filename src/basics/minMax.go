package basics

import (
	"sort"
 	"fmt"
)

//Given five positive integers, find the minimum and maximum values that can be calculated by summing exactly four of the five integers.
// Then print the respective minimum and maximum values as a single line of two space-separated long integers.
func miniMaxSum(a []int) {
    
  sort.Ints(a)
  var cal int
  
  for _, v := range a{
      
    cal += v      
  }

  min := cal - a[len(a)-1]
	max := cal - a[0]
  fmt.Println(min, max)
}