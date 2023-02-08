package main

import "testing"

var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"validTest", 100.0, 10.0, 10.0, false},
	{"validTest", 100.0, 0.0, 0.0, true},
}

func TestNewDivide(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("Expected an error but didnt get one")
			}
		} else {
			if err != nil {
				t.Error("Didnt expect an error but got one")
			}
		}

		if got != tt.expected {
			t.Errorf("Expected %v but got %v", tt.expected, got)
		}
	}

}

func TestDivide(t *testing.T) {
	_, err := divide(10.0, 1.0)
	if err != nil {
		t.Error("Got an error when we shouldnot be getting one")
	}

}

func TestBadDivide(t *testing.T) {
	_, err := divide(10.0, 0.0)
	if err == nil {
		t.Error("Expected an error but didnot get one")
	}
}
