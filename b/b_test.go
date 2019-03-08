package b_test

import (
	"fmt"
	"github.com/go-modules-by-example-staging/cyclic/a"
	"testing"
)

func TestUsingA(t *testing.T) {
	fmt.Printf("Here is A: %v\n", a.AName)
}
