package generic

import (
	"testing"
)

func Test_SetInterSect(t *testing.T) {
	set1 := NewSet(1, 2, 3)
	set2 := NewSet(3, 4, 5)
	res := set1.Intersect(set2)
	if len(res) != 1 {
		t.Errorf("want 1, got %d", len(res))
	}
}

func Test_Contains(t *testing.T) {
	set := NewSet(1, 2, 3)
	assertFalse(t, set.Contains(0))
	assertTrue(t, set.Contains(1))
}

func assertFalse(t *testing.T, cond bool) {
	t.Helper()
	if cond {
		t.Error("expected condition to be false")
	}
}

func assertTrue(t *testing.T, cond bool) {
	t.Helper()
	if !cond {
		t.Error("expected condition to be true")
	}
}
