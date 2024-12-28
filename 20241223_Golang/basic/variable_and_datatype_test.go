package basic_test

import (
	"fmt"
	"testing"
)

var appName = "this is idiomatic"

func TestVariable(t *testing.T) {
	name := "this is idiomatic to"
	x, y, z := "idiomatic X", "idiomatic Y", "idiomatic Z"

	fmt.Println(appName, name, x, y, z)
}
