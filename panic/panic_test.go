package main

import "testing"

func TestPanic(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{1, 2, 3},
		{0, 0, 11},
	}
	for _, tt := range tests {
		if actual := add(tt.a, tt.b); actual != tt.c {
			t.Errorf("calcTriangle(%d, %d); "+
				"got %d; expected %d",
				tt.a, tt.b, actual, tt.c)
		}
	}
}
