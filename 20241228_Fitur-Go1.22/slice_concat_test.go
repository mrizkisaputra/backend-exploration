package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestSliceConcat(t *testing.T) {
	s1 := []string{"a", "b", "c"}
	s2 := []string{"x", "y", "z"}

	result := slices.Concat(s1, s2)
	fmt.Printf("%#v", result)
}
