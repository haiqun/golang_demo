package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Next *TreeNode
}

type TNList struct {
	head *TreeNode
}



func main (){
	t := TNList{head:&TreeNode{Val:1,Next:&TreeNode{Val:3,Next:&TreeNode{Val:5}}}}
	t.List()
	t.Insert(&TreeNode{Val:7,Next:nil})
	t.List()
}

//插入
func (p *TNList)Insert(newCatNode *TreeNode){
	//如果链表为空
	if p.head==nil{
		p.head=newCatNode
		p.head.Next=newCatNode
		return
	}
	temp:=p.head
	//找到队列尾部
	for{
		if temp.Next==p.head{
			break
		}
		temp=temp.Next
	}
	//插入
	temp.Next=newCatNode
	newCatNode.Next=p.head
}

func (p *TNList)List(){
	temp:=p.head
	//如果链表为空
	if temp==nil{
		fmt.Println("null")
		return
	}
	//遍历输出链表
	for{
		fmt.Println(temp.Val)
		if temp.Next==p.head || temp.Next == nil{
			break
		}
		temp=temp.Next
	}
	fmt.Println("====================")
}