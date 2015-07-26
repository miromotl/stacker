package main

import (
    "fmt"
    "github.com/miromotl/stacker/stack"
)

func main() {
    var haystack stack.Stack
    
    // Push some things onto the stack
    haystack.Push("hay")
    haystack.Push(15)
    haystack.Push([]string{"pin", "clip", "needle"})
    haystack.Push(81.52)
    
    // Pop the things from the stack and display them
    for {
        item, err := haystack.Pop()
        if err != nil {
            break
        }
        
        fmt.Println(item)
    }
}