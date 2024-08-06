package constructor

import (
	"testing"
)

func TestOr_Combine(t *testing.T) {
	or := &Or{Types: []Node{new(Any)}}
	or1, err := or.Combine(new(Any))
	if err != nil {
		t.Fatal(err)
	}
	or2, err := or1.Combine(new(String))
	if err != nil {
		t.Fatal(err)
	}
	orString := &Or{Types: []Node{new(String)}}
	if !or2.Equals(orString) {
		t.Errorf("expected %v, got %v", orString, or2)
	}
	or3, err := or2.Combine(new(Any))
	if err != nil {
		t.Fatal(err)
	}
	if !or3.Equals(or2) {
		t.Errorf("expected %v, got %v", or2, or3)
	}
	or4, err := or3.Combine(&Array{Type: new(String)})
	if err != nil {
		t.Fatal(err)
	}
	orArray := &Or{Types: []Node{new(String), &Array{Type: new(String)}}}
	if !or4.Equals(orArray) {
		t.Errorf("expected %v, got %v", orArray, or4)
	}
}
