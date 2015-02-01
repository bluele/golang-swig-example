package example_test

import (
	example "github.com/bluele/golang-swig-example"
	"testing"
)

func TestAdd(t *testing.T) {
	v := example.ExampleAdd(1, 2)
	if (1 + 2) != v {
		t.Error("expected value: %v", 1+2)
	}
}

func TestSetGet(t *testing.T) {
	v := 10
	ex := example.NewExample()
	ex.SetVal(v)
	if ex.GetVal() != v {
		t.Error("expected value: %v", v)
	}
}
