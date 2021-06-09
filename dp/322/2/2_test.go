package dp322_2

import (
	"fmt"
	"testing"
)

//迭代
func coinChange(coins []int, amount int) int {
	dp := make([]int,amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := range dp {
		for _,coin := range coins {
			if i - coin < 0 {
				continue
			}
			if dp[i-coin] == -1 {
				continue
			}
			if dp[i-coin] + 1 < dp[i] {
				dp[i] = dp[i-coin] + 1
			}
		}
		if dp[i] == amount + 1 {
			dp[i] = -1
		}
	}
	return dp[amount]
}

func Test2(t *testing.T) {
	coins := []int{186,419,83,408}
	fmt.Println(coinChange(coins,6248))
}
