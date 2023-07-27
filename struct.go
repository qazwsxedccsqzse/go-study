package main

import (
	"fmt"
)

// 基本語法
type Person struct {
	Name string
	Age  int
}

// 匿名欄位
type Human struct {
	Name   string
	Age    int
	Weight int
}

// struct 如果有內嵌自己, 只能用指標
// 原因是初始化的時候, Go 會先算出需要的空間並且給予零值
// 指標就不會造成無法初始化的問題
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

type Student struct {
	Human      // 匿名欄位，那麼預設 Student 就包含了 Human 的所有欄位
	Speciality string
}

func TestStruct() {
	neil := Student{
		Human: Human{Name: "Neil", Age: 20, Weight: 20}, Speciality: "寫程式",
	}
	fmt.Println("My name is ", neil.Name)
}
