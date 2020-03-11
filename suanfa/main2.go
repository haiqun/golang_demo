package main

import "fmt"

// 单项链条翻转

type listnode struct {
	val  int
	next *listnode
}

func reverseList(head *listnode) *listnode {
	var pre *listnode
	cur := head
	for cur != nil {
		tmp := cur.next // 下个节点先存储起来，为了指向下一个值
		cur.next = pre  // 覆盖下一个key
		pre = cur
		cur = tmp
	}
	return pre
}

func main() {
	l1 := listnode{
		val: 1,
		next: &listnode{
			val: 2,
			next: &listnode{
				val: 3,
				next: &listnode{
					val: 4,
					next: &listnode{
						val: 5,
						next: &listnode{
							val:  6,
							next: nil,
						},
					},
				},
			},
		},
	}
	ret := reverseList(&l1)
	for ret != nil {
		fmt.Printf("%#v->", ret.val)
		ret = ret.next
	}

}
