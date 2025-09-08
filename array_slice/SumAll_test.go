package main

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{3, 4, 5})
	want := []int{3, 12}

	// reflect.DeepEqual 并非类型安全的
	// want := "bob"  // 可以通过编译

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
