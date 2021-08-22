package main

import (
	"log"
	"reflect"
)

// main this app only work for go 1.17 or later
// untuk menjalankan aplikasi ini kamu harus build terlebih dahulu
// $ go build -gcflags=-G=3
func main() {
	// string
	result1 := add[string]("1","2")
	log.Printf("result 1: %s, type: %v", result1, reflect.TypeOf(result1))

	// int
	result2 := add[int](1,2)
	log.Printf("result 1: %d, type: %v", result2, reflect.TypeOf(result2))
}

type Addable interface {
	type int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64, uintptr,
		float32, float64, complex64, complex128,
		string
}

func add[T Addable](a, b T) T {
    return a + b
}