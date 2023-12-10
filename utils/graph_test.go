package utils

import "testing"

func TestGetNeighbours(t *testing.T) {
	//order: always k i j (node kn^2+in+j)
	a := 0
	exp := []int{4, 2, 1}
	act := GetNeighbours(a, 2)
	if len(exp) != len(act) {
		t.Errorf("Array length mismatch: exp: %d, act: %d", len(exp), len(act))
	}

	for i := range exp {
		if exp[i] != act[i] {
			t.Errorf("Position %d: values mismatch. Exp: %d, act: %d", i, exp[i], act[i])
		}
	}
}
