package main

import "fmt"

type Fish struct {
}

func (f Fish) Swim() {
	fmt.Println("我是魚，我會游泳")
}

type FakeFish Fish

func (f FakeFish) FakeSwim() {
	fmt.Println("我是 FakeFish，FakeSwim")
}

type StrongFakeFish Fish

func (f StrongFakeFish) Swim() {
	fmt.Println("我是 StrongFakeFish，FakeSwim")
}

func TestType() {
	fake := FakeFish{}
	// 下面這句會編譯錯誤
	//fake.Swim()
	fake.FakeSwim()

	// 把假魚轉成魚
	converFake := Fish(fake)
	converFake.Swim()

	sFake := StrongFakeFish{}
	// 這邊是呼叫 StrongFakeFish 自己的 swim
	sFake.Swim()
	convertSFake := Fish(sFake)
	convertSFake.Swim()

}
