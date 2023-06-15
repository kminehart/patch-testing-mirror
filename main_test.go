package main

import "testing"

func TestAdd(t *testing.T) {
	if res := Add(1, 2); res != 3 {
		t.Fatalf("Add(1, 2) should equal '%d' but equals '%d'", 3, res)
	}
	if res := Add(2, 2); res != 4 {
		t.Fatalf("Add(2, 2) should equal '%d' but equals '%d'", 3, res)
	}
}
