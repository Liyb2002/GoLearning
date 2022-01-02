package main
import "fmt"

type treeNode struct{
	value int
	left, right *treeNode
}
 
//局部变量不会挂
func createTreeNode(value int) *treeNode{
	return &treeNode{value:value}
}

func (node treeNode)print(){
	fmt.Println(node.value)
}

func main(){ 
	root := treeNode{3, nil, nil}
	root.left=&treeNode{5, nil, nil}
	fmt.Println(root)
	createTreeNode(2)
}