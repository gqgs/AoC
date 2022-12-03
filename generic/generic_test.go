package generic

import "testing"

func Test_Min(t *testing.T) {
	if got := Min(0, 10); got != 0 {
		t.Errorf("want 0, got %d", got)
	}
	if got := Min([]int{0, 1, 2, 3, 4, 5}...); got != 0 {
		t.Errorf("want 0, got %d", got)
	}
	if got := Min([]int{5, 4, 3, 2, 1, 0}...); got != 0 {
		t.Errorf("want 0, got %d", got)
	}
	if got := Min([]uint{0, 1, 2, 3, 4, 5}...); got != 0 {
		t.Errorf("want 0, got %d", got)
	}
}

func Test_Max(t *testing.T) {
	if got := Max(0, 10); got != 10 {
		t.Errorf("want 10, got %d", got)
	}
	if got := Max([]int{0, 1, 2, 3, 4, 5}...); got != 5 {
		t.Errorf("want 5, got %d", got)
	}
	if got := Max([]int{5, 4, 3, 2, 1, 0}...); got != 5 {
		t.Errorf("want 5, got %d", got)
	}
	if got := Max([]uint{0, 1, 2, 3, 4, 5}...); got != 5 {
		t.Errorf("want 5, got %d", got)
	}
}

func Test_SetInterSect(t *testing.T) {
	set1 := NewSet(1, 2, 3)
	set2 := NewSet(3, 4, 5)
	res := set1.Intersect(set2)
	if len(res) != 1 {
		t.Errorf("want 1, got %d", len(res))
	}
}
