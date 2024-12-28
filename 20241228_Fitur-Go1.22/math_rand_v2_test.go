package main

import (
	"fmt"
	"math/rand/v2"
	"testing"
)

var msgs = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
}

func TestMathRandV2(t *testing.T) {
	fmt.Println(rand.Int())
	fmt.Println(rand.Int64())
	fmt.Println(rand.N(100))

	fmt.Println(msgs[rand.N(len(msgs)-1)])
}
