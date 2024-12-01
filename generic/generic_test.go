package generic

import "testing"

func Test_SetInterSect(t *testing.T) {
	set1 := NewSet(1, 2, 3)
	set2 := NewSet(3, 4, 5)
	res := set1.Intersect(set2)
	if len(res) != 1 {
		t.Errorf("want 1, got %d", len(res))
	}
}
