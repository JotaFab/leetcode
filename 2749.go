package main

import (
	"math/bits"
)

func makeTheIntegerZero(num1 int, num2 int) int {
    for i:=0;i<=60; i++{
        s:= int64(num1) - int64(i) * int64(num2)
        if s < 0 {continue}
        if s<int64(i) { continue }
        ones := bits.OnesCount64(uint64(s))
        if ones <= i { return i}
        

    }
    return -1
}