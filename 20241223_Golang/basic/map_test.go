package basic_test

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	var heroes = map[string][]string{
		"tank":     []string{"akai", "balmond", "hilda"},
		"marksman": []string{"irithel", "miya", "layla"},
		"mage":     []string{"edora", "harley", "gord"},
	}

	delete(heroes, "mage")

	for key, value := range heroes {
		fmt.Println(key, ": ", value)
	}
}
