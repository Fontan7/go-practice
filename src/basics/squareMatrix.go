package basics

/*
Given a square matrix, calculate the absolute difference between the sums of its diagonals.
For example, the square matrix arr is shown below:
1 2 3
4 5 6
9 8 9  
The left-to-right diagonal =1+5+9=15 . The right to left diagonal =3+5+9=17 . Their absolute difference is 15-17=2.

Complete the 'diagonalDifference' function below.

The function is expected to return an INTEGER.
The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func DiagonalDifference(arr [][]int) int {
    var leftToR int
    var rightToL int
    
    for i:=0; i < len(arr); i++{
       c := (len(arr[i])-i)-1
       
       leftToR += arr[i][i]
       rightToL += arr[i][c]
    }
    if leftToR > rightToL{
        r := leftToR - rightToL
        return r
    }
    
    return rightToL - leftToR
}