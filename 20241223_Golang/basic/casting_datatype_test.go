package basic_test

import (
	"fmt"
	"strconv"
	"testing"
)

func TestRealNumber(t *testing.T) {
	var valueInt32 int32 = 32
	var valueInt64 int64 = 64

	int32Toint64 := int64(valueInt32)
	int64Toint32 := int32(valueInt64)
	fmt.Println(int32Toint64)
	fmt.Println(int64Toint32)
}

func TestDecimalNumber(t *testing.T) {
	var valueFloat32 float32 = 32.0
	var valueFloat64 float64 = 64.0

	float32Tofloat64 := float64(valueFloat32)
	float64Tofloat32 := float32(valueFloat64)
	fmt.Println(float32Tofloat64)
	fmt.Println(float64Tofloat32)
}

func TestFloatToInt(t *testing.T) {
	var valueFloat float32 = 62.45
	var valueInteger int32 = int32(valueFloat)
	fmt.Println(valueInteger)
}

func TestIntegerToFloat(t *testing.T) {
	var valueInt int32 = 62
	var valueFloat float32 = float32(valueInt)
	fmt.Println(valueFloat)
}

func TestString(t *testing.T) {
	// numeric to string
	var valueInt = 32
	var valueFloat = 64.00
	var intToString = strconv.Itoa(valueInt)
	var floatToString = strconv.FormatFloat(valueFloat, 'f', 2, 64)
	fmt.Println(intToString)
	fmt.Println(floatToString)
}
