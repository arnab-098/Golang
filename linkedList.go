package main

import (
	"fmt"
	"log"
)

type Node struct {
	value int
	next  *Node
}

type List struct {
	head *Node
	tail *Node
}

func createList() List {
	l := List{head: nil, tail: nil}
	return l
}

func isEmptyList(l List) bool {
	return l.head == nil
}

func displayList(l List) {
	if isEmptyList(l) {
		log.Println("List in empty")
	}
	for curr := l.head; curr != nil; curr = curr.next {
		fmt.Printf("%d\t", curr.value)
	}
	fmt.Println()
}

func insertAtFront(num int, l *List) {
	new := &Node{value: num, next: l.head}
	if isEmptyList(*l) {
		l.tail = new
	}
	l.head = new
}

func insertAtEnd(num int, l *List) {
	new := &Node{value: num, next: nil}
	if isEmptyList(*l) {
		l.head = new
	} else {
		l.tail.next = new
	}
	l.tail = new
}

func deleteFromFront(l *List) (int, error) {
	if isEmptyList(*l) {
		return 0, fmt.Errorf("List is empty")
	}
	num := l.head.value
	l.head = l.head.next
	if l.head == nil {
		l.tail = nil
	}
	return num, nil
}

func deleteFromEnd(l *List) (int, error) {
	if isEmptyList(*l) {
		return 0, fmt.Errorf("List is empty")
	}
	num := l.tail.value
	if l.head.next == nil {
		l.head = nil
		l.tail = nil
		return num, nil
	}
	curr := l.head
	for ; curr.next != l.tail; curr = curr.next {
		continue
	}
	curr.next = nil
	l.tail = curr
	return num, nil
}
