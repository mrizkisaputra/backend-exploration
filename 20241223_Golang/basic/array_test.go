package basic_test

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	var names = make([]string, 4)
	names[0] = "muhammat"
	names[1] = "rizki"
	names[2] = "saputra"
	names[3] = "kiki"
	fmt.Println(names)

	var fruits = [...]string{"apple", "banana", "orange"}
	fmt.Println(fruits)

	var matrix = [2][3]int32{
		[3]int32{10, 20, 30},
		[3]int32{40, 50, 60},
	}

	fmt.Println(matrix)
}
