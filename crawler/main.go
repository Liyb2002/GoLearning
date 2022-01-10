package main

import (
	"fmt"
	"regexp"
	"crawler/fetcher"
//	"time"
//	"crawler/worker"
)

func main(){
	all, _ := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	cityUrls:= printCityList(all)

	jobs :=make(chan string, len(cityUrls))
	done := make(chan bool)

	for w:=0; w<3; w++{
		go worker(w, jobs, done)
	}

	for j:=0; j<len(cityUrls); j++{
		jobs <- cityUrls[j]
	}

	close(jobs)
	<- done
	//thisCity, _ := fetcher.Fetch(cityUrls[1])
	//fmt.Println(string(thisCity))
}

func printCityList(contents [] byte) []string{
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"`)
	match := re.FindAllSubmatch(contents, -1)

	cityUrls := make([]string, len(match))

	for i:=0; i<len(match); i++{
		cityUrls[i] = string(match[i][1])
	}

	return cityUrls
}

func worker(id int, jobs <-chan string, done chan<- bool ) {
    for {
		j, more:= <- jobs
		if more{
			fmt.Println("worker", id, "started  job", j)
		}else{
			fmt.Println("worker", id,"received all!")
			done <- true
			return
		}
    }
}