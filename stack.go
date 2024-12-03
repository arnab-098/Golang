package main

import "fmt"

type Stack struct {
	top, size int
	value     []int
}

func createStack(length int) Stack {
	s := Stack{top: -1, size: length, value: make([]int, length)}
	return s
}

func isEmptyStack(s Stack) bool {
	return s.top == -1
}

func isFullStack(s Stack) bool {
	return s.top == s.size-1
}

func push(num int, s *Stack) error {
	if isFullStack(*s) {
		return fmt.Errorf("Stack Overflow")
	}
	s.top++
	s.value[s.top] = num
	return nil
}

func pop(s *Stack) (int, error) {
	if isEmptyStack(*s) {
		return 0, fmt.Errorf("Stack Underflow")
	}
	num := s.value[s.top]
	s.top--
	return num, nil
}

func displayStack(s Stack) {
	if isEmptyStack(s) {
		fmt.Println("Stack is empty")
		return
	}
	for i := 0; i <= s.top; i++ {
		fmt.Printf("%d\t", s.value[i])
	}
	fmt.Println()
}
