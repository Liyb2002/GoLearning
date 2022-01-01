package main

import "fmt"
import "math"

func eval(a,b int, op string)int{
	switch op{
	case "+":
		return a+b
	case "-":
		return a-b
	case "*":
		return a*b
	case "/":
		return a/b
	default:
		panic("unsupportted"+ op)
	}
}

func div(a,b int)(q,r int){
	q=a/b
	r=a%b
	return 

}

func apply(op func(int, int)int, a,b int)int{
	return op(a,b)
}
//函数式编程
func main(){
	fmt.Println(eval(3,4, "+"))
	fmt.Println(apply(
		func (a int, b int)int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3,4 ))
}