package P0001_TwoSum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	assert.Equal(t, []int{0, 1}, twoSum([]int{2, 7, 11, 15}, 9))
	assert.Equal(t, []int{1, 2}, twoSum([]int{3, 2, 4}, 6))
	assert.Equal(t, []int{0, 1}, twoSum([]int{3, 3}, 6))
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if v, ok := m[complement]; ok && v != i {
			return []int{i, m[complement]}
		}
	}
	return []int{0, 0}
}
