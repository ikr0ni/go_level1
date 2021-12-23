package fib

import (
	"fmt"
	"testing"
)

func TestFib(t *testing.T) {
	var f Fib
	f.dig = 10
	expected := 55 //{0:0, 1:1, 2:1, 3:2, 4:3, 5:5, 6:8, 7:13, 8:21, 9:34, 10:55}
	f.fibCalc()
	fmt.Println(f.FibMap)
	fmt.Println(f.iter)
	if f.FibMap[f.iter-1] != expected {
		t.Errorf("Expected %d got %d", expected, f.FibMap[f.iter])
	}
}
