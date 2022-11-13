package s2

import (
	"math/rand"
	"time"
)

type Tree struct {
	Val   int
	Left  *Tree
	Right *Tree
}

func SortedArrayToBST(nums []int) *Tree {
	rand.Seed(time.Now().UnixNano())
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, left, right int) *Tree {
	if left > right {
		return nil
	}

	mid := (left + right + rand.Intn(2)) / 2
	root := &Tree{Val: nums[mid]}
	root.Left = helper(nums, left, mid-1)
	root.Right = helper(nums, mid+1, right)
	return root
}
