package main

import (
	"fmt"
	"regexp"
	"crawler/fetcher"
//	"time"
	"crawler/worker"
)

func main(){
	all, _ := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	cityUrls:= printCityList(all)

	jobs :=make(chan string, len(cityUrls))
	result :=make(chan string, len(cityUrls))
	done := make(chan bool)

	for w:=0; w<3; w++{
		go worker.Work(w, jobs, result, done)
	}

	for j:=0; j<len(cityUrls); j++{
		jobs <- cityUrls[j]
	}

	close(jobs)
	<- done
	//thisCity, _ := fetcher.Fetch(cityUrls[1])
	//fmt.Println(string(thisCity))

	for i:=0; i<len(cityUrls); i++{
		fmt.Println(<-result)
	}
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
