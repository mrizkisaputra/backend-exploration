package basic_test

import (
	"fmt"
	"testing"
)

func TestMakeSliceFromArray(t *testing.T) {

	var arrayNumbers = [...]int32{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	var sliceNumbers = arrayNumbers[2:9]
	sliceNumbers[0] = 200

	var result = append(sliceNumbers, 5000)

	fmt.Println(arrayNumbers)
	fmt.Println(sliceNumbers)
	fmt.Println(result)
}

func TestMakeSlice(t *testing.T) {
	var arrayNumbers = []int32{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	var slices = make([]int32, len(arrayNumbers))

	n := copy(slices, arrayNumbers)
	fmt.Println(slices)
	fmt.Println(n)

}
