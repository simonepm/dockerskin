package hello

import (
	"testing"
)

func TestSayHello(t *testing.T) {
	result := SayHello()
	expected := "Hello World"
	if result != expected {
		t.Errorf("Should say '%v'", expected)
	}
}
