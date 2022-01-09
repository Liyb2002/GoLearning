package main

import (
	"fmt"
	"regexp"
	"crawler/fetcher"
)

func main(){
	all, _ := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	cityUrls:= printCityList(all)
	fmt.Println(cityUrls[1])

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