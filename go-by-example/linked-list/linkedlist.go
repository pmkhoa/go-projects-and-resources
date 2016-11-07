package linkedlist

import "errors"

type Node struct {
    value int
    next *Node
}

type List struct {
    head *Node
    tail *Node
    len int
}

func (list *List) Size() int {
    return list.len
}

func (list *List) Add(value int) {
    node := new(Node)
    node.value = value
    if list.head == nil {
        list.head= node
    } else {
        list.tail.next = node
    }
    list.tail = node
    list.len++
}

func (list *List) Delete(value int) error {
    prev := new(Node)
    for node := list.head; node != nil; node = node.next {
        if node.value == value {
            if node == list.head {
                list.head = node.next
            }
            prev.next = node.next
            list.len--
            return nil
        }
        prev = node
    }
    return errors.New("can't find that value")
}

func (list *List) Search(value int) *Node {
    for node := list.head; node != nil; node = node.next {
        if node.value == value {
            return node
        }
    }
    return nil
}
