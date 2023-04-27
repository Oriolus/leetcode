package main

// https://leetcode.com/problems/smallest-number-in-infinite-set/

import (
	"fmt"
)

/*
храним два объекта
	- минимальный элемент неприрывной последовательности (MinInf)
	- добавленные элементы, которые меньше MinInf нужно где-то хранить. Где
		-- куча					( добавить: O(log(n)), удалить O(log(n)) )
		-- связный список		( добавить: O(n), удалить O(1) )

	k добавлений и удалений (k == n)!

	k * log(n) + k * log(n)  ~	n * log(n) + n * log(n)
	k * n + k				 ~ 	n * n + n
*/

type LiknedList struct {
	Head	*Node
}

func NewLinkedList() LiknedList {
	return LiknedList{}
}

func (this LiknedList) String() string {
	if this.Head == nil {
		return ""
	}

	result := ""
	for cur := this.Head ; cur != nil ; cur = cur.Next {
		result = result + strconv.Itoa(cur.Val) + " -> "
	}
	result += "nil"
	return result
}

func (this *LiknedList) Pop() (int, bool) {
	if this.Head == nil {
		return 0, false
	}

	retVal := this.Head.Val
	this.Head = this.Head.Next

	return retVal, true
}

func (this *LiknedList) Add(val int) {
	if this.Head == nil {
		node := NewNode(val, nil)
		this.Head = &node
		return
	}

	if this.Head.Val == val {
		return
	}

	if this.Head.Val > val {
		newHead := NewNode(val, this.Head)
		this.Head = &newHead
		return
	}

	var prev *Node = nil
	for cur := this.Head ; cur != nil ; cur = cur.Next {
		
		if cur.Val == val {
			return
		}

		if cur.Val > val {
			newNode := NewNode(val, cur)
			prev.Next = &newNode
			return
		}

		prev = cur
	}

	newNode := NewNode(val, nil)
	prev.Next = &newNode
}

type Node struct {
	Val 	int
	Next	*Node
}

func NewNode(val int, next *Node) Node {
	return Node{
		Val: val,
		Next: next,
	}
}

/*
храним два объекта
	- минимальный элемент неприрывной последовательности (MinInf)
	- добавленные элементы, которые меньше MinInf нужно где-то хранить. Где
		-- куча					( добавить: O(log(n)), удалить O(log(n)) )
		-- связный список		( добавить: O(n), удалить O(1) )

	k добавлений и удалений (k == n)!

	k * log(n) + k * log(n)  ~	n * log(n) + n * log(n)
	k * n + k				 ~ 	n * n + n
*/

type SmallestInfiniteSet struct {
	SmallestInf		int
	PreviouseNums	*LiknedList
}

func Constructor() SmallestInfiniteSet {
    return SmallestInfiniteSet{
		SmallestInf: 1,
		PreviouseNums: &LiknedList{},
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	prev, exists := this.PreviouseNums.Pop()
	if exists == true {
		return prev
	}
	this.SmallestInf += 1
	return this.SmallestInf - 1
}

func (this *SmallestInfiniteSet) AddBack(num int)  {
	if num >= this.SmallestInf {
		return
	}

	this.PreviouseNums.Add(num)
}


/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */


func main() {

	set := Constructor()
	set.AddBack(2)
	fmt.Println(set.PopSmallest())
	fmt.Println(set.PopSmallest())
	fmt.Println(set.PopSmallest())
	set.AddBack(1)
	fmt.Println(set.PopSmallest())
	fmt.Println(set.PopSmallest())
	fmt.Println(set.PopSmallest())
}