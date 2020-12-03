package simple_factory

import "testing"

func TestType1(t *testing.T) {
	api := NewAPI(1)
	s := api.Say("Shou")
	if s != "Hi, Shou" {
		t.Fatal("Type1 test fail")
	}
}

func TestType2(t *testing.T) {
	api := NewAPI(2)
	s := api.Say("Shou")
	if s != "Hello, Shou" {
		t.Fatal("Type2 test fail")
	}
}
