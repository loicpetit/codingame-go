package test

import (
	"fmt"
	"testing"
)

func TestStdoutInterceptor(t *testing.T) {
	interceptor := NewStdoutInterceptor()
	interceptor.Intercept()
	fmt.Println("1 2 5")
	var i1, i2, i3 int
	interceptor.Scan(&i1, &i2, &i3)
	if i1 != 1 {
		t.Errorf("i1 should be 1 bu t is %d", i1)
	}
	if i2 != 2 {
		t.Errorf("i2 should be 2 bu t is %d", i2)
	}
	if i3 != 5 {
		t.Errorf("i3 should be 5 bu t is %d", i3)
	}
	interceptor.Close()
	fmt.Println("Should be print in stdout")
}
