package main

import (
    "fmt"
    "math"
)

type Tree struct {
    Value int
    Left *Tree
    Right *Tree
}

func getMax(t *Tree) int{
    if t == nil {
        return -1
    }
    leftMax := getMax(t.Left)
    rightMax := getMax(t.Right)
    yz := math.Max(float64(leftMax), float64(rightMax))
    return int(math.Max(float64(t.Value), yz))
}

func main(){
    t := &Tree{
        Value:  3,
        Left:  &Tree{
            Value:  1,
            Left:  nil,
            Right: nil,
        },
        Right: &Tree{
            Value: 22,
            Left:  nil,
            Right: nil,
        },
    }

    fmt.Println(getMax(t))

}
