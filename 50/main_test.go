package leetcoded

import "testing"
import "github.com/stretchr/testify/assert"

func Test1(t *testing.T) {
	input := struct {
		x float64
		n int
	}{2.00000, 10}
	out := myPow(input.x, input.n)
	assert.Equal(t, 1024.00000, out)
}
