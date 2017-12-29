package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	grid := [][]byte{
		[]byte{0x31, 0x31, 0x31, 0x31, 0x30},
		[]byte{0x31, 0x31, 0x30, 0x31, 0x30},
		[]byte{0x31, 0x31, 0x30, 0x30, 0x30},
		[]byte{0x30, 0x30, 0x30, 0x30, 0x30},
	}
	out := numIslands(grid)
	assert.Equal(t, 1, out)

}

func Test2(t *testing.T) {
	grid := [][]byte{
		[]byte("11000"),
		[]byte("11000"),
		[]byte("00100"),
		[]byte("00011"),
	}
	out := numIslands(grid)
	assert.Equal(t, 3, out)
}
