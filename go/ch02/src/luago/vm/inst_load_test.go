package vm

import (
	"fmt"
	"testing"
)

func TestLoadk(t *testing.T)  {
	var i Instruction = (2 + (12 << 8)) << 6;
	a, bx := i.ABx()
	fmt.Println(a, bx)
}
