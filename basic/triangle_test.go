package main

import "testing"

func TestTriangle(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{12, 35, 37},
	}
	for _, tt := range tests {
		if actual := calcTriangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("calcTriangle(%d, %d); got %d; expected %d", tt.a, tt.b, actual, tt.c)
		}
	}
}

func BenchmarkTriangle(b *testing.B) {
	ta, tb := 30000, 40000
	ans := 50000

	b.Logf("start to bench test")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if actual := calcTriangle(ta, tb); actual != ans {
			b.Errorf("calcTriangle(%d, %d); got %d; expected %d", ta, tb, actual, ans)
		}
	}
}
