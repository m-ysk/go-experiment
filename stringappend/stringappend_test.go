package main

import (
	"testing"
)

func StringAdd() {
	var s string
	for i := 0; i < 100000; i++ {
		s += "a"
	}
}

func Benchmark_StringAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringAdd()
	}
}

func StringAppend() {
	var b []byte
	for i := 0; i < 100000; i++ {
		b = append(b, "a"...)
	}
}

func Benchmark_StringAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringAppend()
	}
}

func StringAppend2() {
	b := make([]byte, 0, 100000)
	for i := 0; i < 100000; i++ {
		b = append(b, "a"...)
	}
}

func Benchmark_StringAppend2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringAppend2()
	}
}
