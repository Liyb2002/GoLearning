package worker

import (
//	"crawler/fetcher"
//	"fmt"
)

func Work(id int, jobs <-chan string, result chan<- string, done chan<- bool ) {
    for {
		j, more:= <- jobs
		if more{
	//		fmt.Println("worker", id, "started  job", j)
			result <- j
		}else{
	//		fmt.Println("worker", id,"received all!")
			done <- true
			return
		}
    }
}