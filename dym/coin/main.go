package main

import (
    "fmt"
    "math"
)

// 给你 k 种面值的硬币，面值分别为 c1, c2 ... ck，每种硬币的数量无限，再给一个总金额 amount，问你最少需要几枚硬币凑出这个金额，如果不可能凑出，算法返回 -1
func coinsConnect(coins []int, amount int) int {

    // 动态规划
    var dp func(int) int
    dp = func (total int) int {
        // 确定base
        if total == 0 {//硬币找完了
            return 0
        }
        if total < 0 {
            return -1
        }
        // 求最小值 取正无穷
        res := math.Inf(1)
        for _,coin := range coins {
            subp := dp(total - coin)
            if subp == - 1 {//子问题无解，跳过
                continue
            }
            res = math.Min(res, float64(subp) + 1) //找到一枚硬币
        }
        if !math.IsInf(res, 1) {//非正无穷，返回结果
            return int(res)
        }
        return -1
    }

    return dp(amount)
}

func main() {
    fmt.Println(coinsConnect([]int{2}, 6))
}
