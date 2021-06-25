package dp300_1

import (
	"fmt"
	"testing"
)


func lengthOfLIS(nums []int) int {
	dp := make([]int,len(nums))
	maxLen := 1
	for i:= range nums {
		dp[i] = 1
		for j:= 0; j < i; j++ {
			fmt.Printf("nums[%d]:%d,nums[%d]:%d\n",j,nums[j],i,nums[i])
			if nums[j] < nums[i] {
				dp[i] = max(dp[i],dp[j]+1)
			}
			fmt.Printf("dp[%d]:%d,dp[%d]:%d\n",i,dp[i],j,dp[j])
		}
		maxLen = max(maxLen,dp[i])
	}
	return maxLen
}

func max(i,j int) int {
	if i >= j {
		return i
	} else {
		return j
	}
}

func Test1(t *testing.T) {
	nums := []int{4,10,4,3,8,9}
	fmt.Println(lengthOfLIS(nums))
}
