package main

import "testing"

func BenchmarkReadJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadJson()
	}
}

func BenchmarkReadCsvFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadCsvFile()
	}
}
