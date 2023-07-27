package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) Append(data int) {
	newNode := &Node{Data: data}

	if l.Head == nil {
		l.Head = newNode
	} else {
		current := l.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
}

// 新增的 Remove 方法
func (l *LinkedList) Remove(data int) {
	if l.Head == nil {
		return
	}

	// 如果头节点就是要移除的节点
	if l.Head.Data == data {
		l.Head = l.Head.Next
		return
	}

	// 找到要移除的节点和其前一个节点
	current := l.Head
	for current.Next != nil && current.Next.Data != data {
		current = current.Next
	}

	// 如果找到，移除节点
	if current.Next != nil {
		current.Next = current.Next.Next
	}
}

// 使用 pointer to pointer 移除 node
func (l *LinkedList) Remove2(data int) {
	indirect := &l.Head

	for (*indirect).Data != data {
		indirect = &(*indirect).Next
	}

	*indirect = (*indirect).Next
}

func (l *LinkedList) Display() {
	node := l.Head
	for node != nil {
		fmt.Println(node.Data)
		node = node.Next
	}
}

// func main() {
// 	linked := &LinkedList{}
// 	linked.Append(1)
// 	linked.Append(2)
// 	linked.Append(3)

// 	linked.Display()

// 	fmt.Println("After removing element:")
// 	linked.Remove(2)
// 	linked.Display()
// }
