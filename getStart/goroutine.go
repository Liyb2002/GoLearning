package main

import ("fmt" 
"time")

func main(){
	for i:=0; i<1000; i++{
		//go 使得并发执行
		go func (i int){
			for {
			fmt.Printf("hello from goroutine %d \n", i)
		}
		}(i)
	}
	time.Sleep(time.Millisecond)
}