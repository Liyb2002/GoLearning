package worker

import (
//	"crawler/fetcher"
	"fmt"
)

func Work(id int, jobs <- chan string){
	fmt.Println("worker % is working!", id)
		url:= <- jobs
		fmt.Println("worker %d working on job %s", id, url)

	

}