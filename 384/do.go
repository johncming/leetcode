package leetcode

import (
	crand "crypto/rand"
	"log"
	"math"
	"math/big"
	"math/rand"
)

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	s := Solution{nums: nums}

	n, err := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		panic(err)
	}

	rand.Seed(n.Int64())

	return s
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.nums
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	var newNums []int

	size := len(this.nums)

	index := rand.Intn(size)

	log.Println(index)

	newNums = append(newNums, this.nums[:index]...)
	newNums = append(newNums, this.nums[index:size]...)

	return newNums
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
