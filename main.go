package main

import (
	"fmt"
	"github.com/nvkv/halp/pkg/types/v1"
)

func main() {
	b := types.Meal{
		Name: "Test Meal",
	}

	fmt.Printf("Halp! I need somebody! %#v", b)
}
