package basics

import (
 	"fmt"
)

func fizzBuzz(n int32) {
    var i int32
    for i=1; i <= n; i++{
        
		var output string

     	if i%3 == 0 && i%5 != 0{
            output += "Fizz"
        }else if i%3 != 0 && i%5 == 0{
            output += "Buzz"
        } else {
			output = string(i)
		}

		fmt.Println(output)
	}
}