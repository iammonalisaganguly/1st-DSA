package main

import "fmt"

func main() {
	StackObj := NewStackConstructor(5)
	StackObj.push(30)
	StackObj.push(90)
	StackObj.push(40)
	StackObj.push(30)
	StackObj.push(90)
	StackObj.push(40)
	StackObj.pop()
	StackObj.pop()
}

type StackClass struct {
	arr  []int
	top  int
	size int
}

func NewStackConstructor(size int) *StackClass {
	return &StackClass{
		size: size,
		arr:  make([]int, size),
		top:  -1,
	}
}

func (s *StackClass) push(eleme int) {
	if s.top == s.size-1 { // ami check korchi ata ki full ache .
		fmt.Println("Stack is full")
	} else {
		s.top++
		s.arr[s.top] = eleme

	}
}

func (s *StackClass) pop() {
	if s.top == -1 {
		fmt.Println("Stack is empty")
	} else {
		s.top--
	}

}
