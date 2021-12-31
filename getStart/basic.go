package main
import "fmt"

var(
	aa=3
	bb="k"
	cc=true
)
func variableInitValue(){
	var a int =3
	var s string = "abc"
	fmt.Println("hh")
	fmt.Printf("%d %q \n", a,s)
}
func enums(){
	const(
		cpp = iota
		_
		java
		python
	)
	fmt.Println(cpp,java,python)
}
func variableShorter(){
	a,b,c,d := 3,4,"thh",true
	b=5
	fmt.Println("variableShorter", a,b,c,d)
}

func main(){
	fmt.Println("hello")
	variableShorter()
	fmt.Println(aa,bb,cc)
	enums()
}