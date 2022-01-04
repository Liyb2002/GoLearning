package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int) {
    fmt.Printf("Worker %d starting\n", id)

    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {

    var wg sync.WaitGroup

    for i := 1; i <= 5; i++ {
        wg.Add(1)

        i := i

        go func() {
			//wg.Done() decrements the waitgroup counter by one
            defer wg.Done()
            worker(i)
        }()
    }

	//blocks until the waitgroup counter is zero
    wg.Wait()
	fmt.Println("all is done")

}