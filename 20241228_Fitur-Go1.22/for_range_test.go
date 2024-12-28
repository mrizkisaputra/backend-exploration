package main

import (
	"fmt"
	"testing"
)

func TestForRangeInteger(t *testing.T) {
	for i := range 10 {
		fmt.Println(i)
	}
}
