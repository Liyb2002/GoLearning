package main
import (
	"regexp"
	"fmt"
)

const text = "my email is yl4092@Columbia.edu  yl40492@Columbia.edu  yl2@Columbia.edu"

func main(){
	re := regexp.MustCompile(`[a-zA-Z0-9]+@.+\..+`)

	match := re.FindAllString(text, -1)
	fmt.Println(match)
}