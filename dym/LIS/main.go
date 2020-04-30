package main

import "fmt"

// 最长子序列: 序列不是连续的 通过dp数组迭代，由前面的最长长度推断当前位的长度

func dp(nums []int) int {
    list := make([]int, len(nums))
    for i,_ := range nums {
        list[i] = 1
    }

    for i:=0;i < len(nums);i++ {
        for j:=0;j<i;j++ {
            if nums[i] > nums[j] {// 当前位与前长度比较，若找到更大的数
            // dp数组当前位赋值
                list[i] = Max(list[i], list[j] + 1)
            }
        }
    }

    res := 0
    for i:=0;i<len(list);i++ {
        res = Max(res, list[i])
    }

    return res
}

func Max(a ,b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {

    fmt.Println(dp([]int{1,2,3,14,5}))

}
