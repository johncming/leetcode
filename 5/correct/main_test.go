package leetcode

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	s := "babad"
	out := longestPalindrome(s)
	fmt.Printf("%+v\n", out) // output for debug
}

func Test2(t *testing.T) {
	s := "a"
	out := longestPalindrome(s)
	fmt.Printf("%+v\n", out) // output for debug
}

func Test3(t *testing.T) {
	s := "babadada"
	out := longestPalindrome(s)
	fmt.Printf("%+v\n", out) // output for debug
}
