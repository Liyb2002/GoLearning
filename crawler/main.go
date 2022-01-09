package main

import (
	"fmt"
	"regexp"
	"crawler/fetcher"
)

func main(){
	fetcher.TestPrint()
	/*
	resp, err := http.Get("http://www.zhenai.com/zhenghun")

	if err !=nil{
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK{
		all, err := ioutil.ReadAll(resp.Body)
		if err!=nil{
			panic(err)
		}
		//fmt.Printf("%s \n", all)
		printCityList(all)
	}
*/

}

func printCityList(contents [] byte){
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"`)
	match := re.FindAllSubmatch(contents, -1)
	fmt.Println(string(match[1][1]))
}