package main

import (
    "fmt"
    "time"
)

// 备忘录
var memo map[int]int

// 递归求解
func fb(num int) int {
    if num == 1 || num == 2 {
        return 1;
    }
    if val,ok := memo[num]; ok {
        return val
    }
    memo[num] = fb(num - 1) + fb(num -2 )
    return memo[num]
}

// 动态规划数组迭代，自底向上求解
func fbDP(num int) int {
    dp := make([]int, num + 1)
    dp[1] = 1
    dp[2] = 1
    for i:= 3;i<=num; i++ {
        dp[i] = dp[i - 1] + dp[i - 2]
    }
    return dp[num]
}

// 状态转移只涉及到前两个变量
func fbDP1(num int) int {
    dp1 := 1
    dp2 := 1
    for i:= 3;i<=num; i++ {
        sum := dp1 + dp2
        dp1 = dp2
        dp2 = sum
    }
    return dp2
}

func main() {

    memo = map[int]int{}

    start := time.Now()
    fmt.Println(fb(50))
    fmt.Println(time.Since(start))

    start1 := time.Now()
    fmt.Println(fbDP(50))
    fmt.Println(time.Since(start1))

    start2 := time.Now()
    fmt.Println(fbDP1(50))
    fmt.Println(time.Since(start2))
}
