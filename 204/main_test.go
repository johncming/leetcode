package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	n := 10
	out := countPrimes(n)
	assert.Equal(t, 4, out)
}

// result: long long
func Benchmark1(b *testing.B) {
	countPrimes(499979)
}

func Benchmark2(b *testing.B) {
	countPrimes(1500000)
}
