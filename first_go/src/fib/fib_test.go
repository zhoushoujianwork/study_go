package fib

import (
	"fmt"
	"testing"
)

var c int

func TestFibList(t *testing.T) {
	// var a int = 1
	// var b int = 1
	c = 5
	// d := 1

	var (
		a int = 1
		b int = 1
	)

	for i := 0; i < 15; i++ {
		// t.Log(b)
		fmt.Println(b)
		tmp := a
		a = b
		b = tmp + a
	}
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	a, b = b, a
	fmt.Println(a, b)
}

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

func TestConstant(t *testing.T) {
	fmt.Println(Monday, Tuesday, Wednesday)
}

const (
	Readable = 1 << iota
	Writeable
	Exchange
)

func TestConstant02(t *testing.T) {
	// 0111 7  true true true
	// 0001 1  true false false
	f := 1
	fmt.Println(f&Readable == Readable, f&Writeable == Writeable, f&Exchange == Exchange)
}
