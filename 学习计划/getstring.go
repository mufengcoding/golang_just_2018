package main

import (
	"fmt"
	"log"
)

//匹配出日志中的{}

type Stack struct {
	size int64
	top  int64
	data []interface{}
}

func MakeStack(size int64) Stack {
	s := Stack{}
	s.size = size
	s.data = make([]interface{}, size)

	return s
}

func (s *Stack) push(e interface{}) bool {

	if s.IsFull() {
		log.Printf("炸了炸了")
		return false
	}

	s.data[s.top] = e
	s.top++
	return true
}

func (s *Stack) Pop() (r interface{}, err error) {

	if s.IsEmpty() {
		err = fmt.Errorf("空栈")
		log.Printf("空栈")
		return
	}
	s.top--
	r = s.data[s.top]
	return
}

//判断栈是否满
func (s *Stack) IsFull() bool {

	return s.top == s.size
}

//判断栈是否为空
func (s *Stack) IsEmpty() bool {

	return s.top == 0
}

func (s *Stack) Traverse(fn func(r interface{}), goorto bool) {

	//go true遍历进栈   false 遍历出栈
	if goorto {
		var i int64 = 0
		for ; i < s.top; i++ {
			fn(s.data[i])

		}

	} else {

		for i := s.top - 1; i >= 0; i-- {
			fn(s.data[i])
		}
	}

}

func (s *Stack) Mate(str string) {

	for _, v := range str {
		s.push(string(v))

		if string(v) == "{" || string(v) == "}" {

			fmt.Println(s.Pop())
			//fmt.Println(id)
		}

	}

}

func main() {
	str := "{adfd{{zjhgh}luy}"
	s := MakeStack(20)
	s.Mate(str)
	fmt.Println(s)

}
