package types

import (
	"fmt"
	"testing"
	"testing/quick"
)

func TestValidation(t *testing.T) {
	gen := func(d Day) bool {
		fmt.Printf("%#v\n", d)
		return true
	}

	if err := quick.Check(gen, nil); err != nil {
		t.Error(err)
	}
}
