package dp322_1

import (
	"fmt"
	"testing"
)

//递归
func coinChange(coins []int, amount int) int {
	memo := make(map[int]int,amount)
	return dp(coins,amount,&memo)
}

func dp (coins []int, amount int,memo *map[int]int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	if (*memo)[amount] != 0 {
		return (*memo)[amount]
	}
	res := amount + 1
	for _,coin := range coins {
		//fmt.Println("coin:",coin)
		subRes := dp(coins,amount-coin,memo)
		if subRes == -1 {
			continue
		}
		if subRes + 1 < res {
			res = subRes + 1
		}
	}
	if res == amount + 1 {
		res = -1
	}
	(*memo)[amount] = res
	//fmt.Println("amount:",amount)
	//fmt.Println("res:",res)
	//fmt.Println("memo:",*memo)
	return res
}


func Test1(t *testing.T) {
	coins := []int{186,419,83,408}
	fmt.Println(coinChange(coins,6248))
}
