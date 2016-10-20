package pack

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	if result != 3 {
		t.Log("The number wasnt three")
		t.Fatal()
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(1, 2)
	if result != 2 {
		t.Log("The number wasnt two")
		t.Fatal()
	}
}
func TestDivide(t *testing.T) {
	result := Divide(4, 2)
	if result != 2 {
		t.Log("The number wasnt two")
		fmt.Println(result)
		t.Fatal()
	}
}
