package leetcode

import "testing"

func Test_intersect(t *testing.T) {
	n1 := intersect([]int{1, 2, 2, 1}, []int{2, 2})
	if len(n1) != 2 {
		t.Fatal("len(n1) != 2")
	}
	if n1[0] != 2 && n1[1] != 2 {
		t.Fatal(n1)
	}

	n2 := intersect([]int{4, 7, 9, 7, 6, 7}, []int{5, 0, 0, 6, 1, 6, 2, 2, 4})
	if len(n2) != 2 {
		t.Fatal("len(n2) != 2")
	}
	if n2[0] != 4 || n2[1] != 6 {
		t.Fatal(n2)
	}
}
