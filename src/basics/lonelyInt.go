package basics

import(
	"go-practice/src/tools"
)
//Recieves an array in which every number occurs twice exept one
//It finds that number, the lonely integer
func Lonelyinteger(a []int) int {
    
	uniqueSum := tools.Sum(tools.Unique(a))
	allSum := tools.Sum(a)

	lonely := uniqueSum*2 - allSum

	return lonely
}