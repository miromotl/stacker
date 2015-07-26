package stack

import "errors"

// Our custom type 'Stack', a slice of empty interfaces
type Stack []interface{}

// Returns the length of the Stack
// It is customary to implement this method
func (stack Stack) Len() int {
    return len(stack)
}

// Returns the capacity of the Stack
// It is customary to implement this method
func (stack Stack) Cap() int {
    return cap(stack)
}

// Returns true if the Stack is empty
func (stack Stack) IsEmpty() bool {
    return len(stack) == 0
}

// Pushes an item onto the Stack
func (stack *Stack) Push(x interface{}) {
    *stack = append(*stack, x)
}

// Returns the top item from the Stack
func (stack Stack) Top() (interface{}, error) {
    if len(stack) == 0 {
        return nil, errors.New("can't Top() an empty stack")
    }
    
    return stack[len(stack)-1], nil
}

// Pops the top item from the Stack
func (stack *Stack) Pop() (interface{}, error) {
    theStack := *stack
    
    if len(theStack) == 0 {
        return nil, errors.New("can't Pop() an empty stack")
    }
    
    x := theStack[len(theStack)-1]
    *stack = theStack[:len(theStack)-1]
    
    return x, nil
}