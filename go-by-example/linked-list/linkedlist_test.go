package linkedlist

import "testing"

func assert( t *testing.T, actual int, expected int) {
    if actual != expected {
        t.Errorf("expected %v, but got %v", expected, actual)
    }
}

func assertNodes( t *testing.T, actual *Node, expected *Node) {
    if actual != expected {
        t.Errorf("expected %v, but got %v", expected, actual)
    }
}

func TestSize(t *testing.T) {
    list := new(List)
    assert(t, list.Size(), 0)

    list.Add(4)
    assert(t, list.Size(), 1)

    list.Add(2)
    assert(t, list.Size(), 2)

    list.Delete(4)
    assert(t, list.Size(), 1)
}

func TestAdd(t *testing.T) {
    list := new(List)
    list.Add(4)
    assert(t, list.Size(), 1)
    assert(t, list.head.value, 4)

    list.Add(2)
    assert(t, list.Size(), 2)
    assert(t, list.tail.value, 2)
}

func TestDelete(t *testing.T) {
    list := new(List)
    list.Add(4)
    list.Delete(4)
    assert(t, list.Size(), 0)
    assertNodes(t, list.head, nil)

    list.Add(2)
    list.Add(5)
    list.Delete(2)
    assert(t, list.Size(), 1)
    assertNodes(t, list.head, list.tail)
}

func TestSearch(t *testing.T) {
    list := new(List)
    list.Add(2)
    list.Add(4)
    list.Add(4)
    list.Add(8)

    actualNode := list.Search(4)
    expected := list.head.next
    assertNodes(t, actualNode, expected)

    assertNodes(t, list.Search(9), nil)
}


