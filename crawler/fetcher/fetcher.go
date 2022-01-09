package fetcher

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func Fetch(url string) ([]byte, error){

	resp, err := http.Get(url)
	if err !=nil{
		panic(err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK{
		all, err := ioutil.ReadAll(resp.Body)
		if err!=nil{
			panic(err)
		}
		//fmt.Printf("%s \n", all)
		return all, err
}

return nil, nil
}

func TestPrint() string{
	fmt.Println("hi")
	return "hi"
}