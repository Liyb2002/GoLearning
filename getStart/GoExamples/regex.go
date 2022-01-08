package main
import (
	"regexp"
	"fmt"
)

const text = "my email is yl4092@Columbia.edu"

func main(){
	re := regexp.MustCompile(`[a-zA-Z0-9]+@.+\..+`)

	match := re.FindString(text)
	fmt.Println(match)
}