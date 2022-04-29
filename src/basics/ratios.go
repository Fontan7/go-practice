package basics

import (
 	"fmt"
)

//Print the ratios of positive, negative and zero values in the array.
// Each value should be printed on a separate line with 6 digits after the decimal. 
//The function should not return a value.
func plusMinus(a []int32) {
    
    base := float32(len(a))
    var pos, neg, z float32
    
    for _, n := range a{
        if n > 0{
            pos += 1
        } else if n < 0{
            neg += 1
        } else{
            z += 1
        }
    }
    
    pos = pos/base
    neg = neg/base
    z = z/base
    
    fmt.Printf("%.6f",pos)
    fmt.Println()
    fmt.Printf("%.6f",neg)
    fmt.Println()
    fmt.Printf("%.6f",z)
}