package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	var x IntSet
	mx := make(map[int]bool)

	// Add
	x.Add(1)
	mx[1] = true

	for i, _ := range mx {
		if !x.Has(i) {
			t.Errorf("invalid result. IntSet x:%v, set mx:%v", x, mx)
		}
	}
}

func TestUnion(t *testing.T) {
	var x, y IntSet

	mx := make(map[int]bool)
	my := make(map[int]bool)

	// Add
	x.Add(1)
	mx[1] = true

	y.Add(2)
	my[2] = true

	// Union
	x.UnionWith(&y)
	merged := merge(mx, my)

	for i, _ := range merged {
		if !x.Has(i) {
			t.Errorf("invalid result. IntSet x:%v, set merged:%v", x, merged)
		}
	}

}

func merge(m1, m2 map[int]bool) map[int]bool {
	ans := map[int]bool{}

	for k, v := range m1 {
		ans[k] = v
	}
	for k, v := range m2 {
		ans[k] = v
	}
	return (ans)
}
