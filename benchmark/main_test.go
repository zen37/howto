package main

import "testing"

func BenchmarkReadJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadJson()
	}
}

func BenchmarkReadJsonOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadJsonOne()
	}
}

func BenchmarkReadJsonBig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadJsonBig()
	}
}

func BenchmarkReadJsonBigProto(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadJsonBigProto()
	}
}

func BenchmarkReadCsvOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadCsvOne()
	}
}

func BenchmarkReadCsv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadCsv()
	}
}

func BenchmarkReadCsvBig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadCsvBig()
	}
}
