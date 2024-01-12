package main

import "fmt"


type List[T any] struct {
    head *Node[T]
    tail *Node[T]
}

type Node[T any] struct {
    next *Node[T]
    val T
}

func (ll *List[T]) Add(v T)  *List[T]{
    newNode := &Node[T]{nil, v}

    if ll.head == nil {
        ll.head = newNode
        ll.tail = newNode
    } else {
        ll.tail.next = newNode
        ll.tail = ll.tail.next
    }

    return ll
}

func DisplayList[T any](ll *List[T]) string {
    var out string = ""
    iter := ll.head
    
    for iter != nil {
        out += fmt.Sprintf("%v ->", iter.val)
        iter = iter.next
    }
    return out
}

func (ll *List[T]) Display() string {
    return DisplayList(ll)
}

func CreateList[T any](arr []T) *List[T] {
    list := &List[T]{}

    for _, v := range arr {
        list.Add(v)
    }
    return list
}

func main() {
    ArrayInt := []int{1, 2, 3, 4, 5, 6, 7}
    list1 := CreateList(ArrayInt)

    StrArr := []string{"foo", "bar", "look"}
    list2 := CreateList(StrArr)

    fmt.Println(list1.Display())
    fmt.Println(list2.Display())
}
