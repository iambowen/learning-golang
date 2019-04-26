package ch2

import (
	"testing"
)
const (
    Monday  = iota + 1
    Tuesday
)
const (
    Open = 1 << iota
    Close
)

func TestFib(t *testing.T) {
    t.Log(Monday, Tuesday)
    t.Log(Open, Close)
}


func TestSomething(t *testing.T) {
    a := 1
    aPtr := &a
    t.Logf("%T %T", a, aPtr)
    t.Log(2&^2)
}
