package main

import "io/ioutil"
import "fmt"

func grade(score int)string{
	g := ""
	switch{
	case score<60:
		g="F"
	case score<80:
		g="C"
	default:
		panic("wrong score")
	}
	return g
}

func main(){
	const filename="hello.go"
	contents, err:=ioutil.ReadFile(filename)
	if err !=nil{
		fmt.Println(err)
	}else{
		fmt.Println(contents)
	}

	fmt.Println(grade(30))
}