package main

import "fmt"

/*
   WAP to implement a stack using array.
       a) Insert the value in the stack using function push
       b) Delete the value in the stack using function pop
       c) Display the stack using function display
*/

type stack struct {
    top int
    array []int
}

var s = stack{
    top: -1,
    array: []int{},
}

func push(num int) {
        s.top = s.top + 1
        s.array = append(s.array, num)
}

func pop() {
    if s.top == -1 {
        fmt.Println("UnderFlow")
        return
    }
    fmt.Printf("%d is deleted from stack\n", s.array[s.top])
    s.array = s.array[:s.top]
    s.top = s.top - 1
}

func display() {
    for i := s.top; i >= 0; i-- {
        fmt.Printf("%v ", s.array[i])
    }
    fmt.Println()
}

func main() {

    push(12)
    push(11)
    push(13)
    display()

}
