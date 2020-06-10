package main

import "testing"

func AllocStack() int {
	// capを8192以上にするとheapにescapeされる
	s := make([]int, 0, 8191)
	for i := 0; i < 8191; i++ {
		s = append(s, i)
	}

	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

func BenchmarkMemAlloc_AllocStack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AllocStack()
	}
}

func AllocHeap() int {
	// capを8192以上にするとheapにescapeされる
	s := make([]int, 0, 8192)
	for i := 0; i < 8192; i++ {
		s = append(s, i)
	}

	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

func BenchmarkMemAlloc_AllocHeap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AllocHeap()
	}
}
