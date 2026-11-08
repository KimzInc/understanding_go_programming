package main

import "testing"

// Package-level sinks. Assigning results to a global forces the compiler to
// treat the value as "used elsewhere", so it can't optimize the work away.
var (
	sinkPerson Person
	sinkPtr    *Person
	sinkInt    int
)

// 1. Return by value. Expect: 0 allocs/op.
func BenchmarkMakePerson(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkPerson = MakePerson("John", "Doe", 30)
	}
}

// 2. Return a pointer AND store it in a global, so the pointer outlives the
// loop iteration. It MUST escape. Expect: 1 alloc/op, with or without inlining.
func BenchmarkMakePersonPtr_Escapes(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkPtr = MakePersonPtr("Jane", "Smith", 25)
	}
}

// 3. Return a pointer but only read from it locally; the pointer never leaves
// this function. With inlining on, escape analysis can keep it on the stack
// (0 allocs/op). With inlining off (-gcflags=-l), MakePersonPtr is a real
// call, so the Person must live on the heap (1 alloc/op).
func BenchmarkMakePersonPtr_Local(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		p := MakePersonPtr("Jane", "Smith", 25)
		sinkInt += p.Age
	}
}