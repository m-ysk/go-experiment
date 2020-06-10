package main

import (
	"fmt"
	"math/rand"
	"testing"
)

const N = 10000000

var (
	s = make([]int, N, N)
	r = make([]int, N, N)
)

func Sequential(n int) int64 {
	var sum int64
	for i := 0; i < n; i++ {
		sum += int64(s[i])
	}
	return sum
}

func BenchmarkLocality_Sequential(b *testing.B) {
	for i := 0; i < N; i++ {
		s[i] = rand.Intn(100)
		r[i] = i
	}

	var sum int64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum += Sequential(N)
	}
	b.StopTimer()

	fmt.Println(sum)
}

func Random(n int) int64 {
	var sum int64
	for i := 0; i < n; i++ {
		sum += int64(s[r[i]])
	}
	return sum
}

func BenchmarkLocality_Random(b *testing.B) {
	for i := 0; i < N; i++ {
		s[i] = rand.Intn(100)
		r[i] = rand.Intn(N)
	}

	var sum int64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum += Random(N)
	}
	b.StopTimer()

	fmt.Println(sum)
}

