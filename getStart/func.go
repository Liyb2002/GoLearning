package main

import "fmt"

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

func main(){
	fmt.Println(eval(3,4, "+"))
}