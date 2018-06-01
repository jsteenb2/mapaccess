package mapaccess_test

import (
	"testing"

	"github.com/jsteenb2/mapaccess"
)

func Benchmark_AccessWithEmptyStruct(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mapaccess.AccessWithEmptyStructs([]int{0, 3, 11, 1, 20, 30, 1000000, 4})
	}
}

func Benchmark_AccessWithBool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mapaccess.AccessWithBool([]int{0, 3, 11, 1, 20, 30, 1000000, 4})
	}
}
