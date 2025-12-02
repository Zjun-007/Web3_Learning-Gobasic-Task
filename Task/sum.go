// 两数之和
// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

// 你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。

// 你可以按任意顺序返回答案。

// 示例：
// 输入：nums = [3,2,4], target = 6
// 输出：[1,2]


package main

import "fmt"

func twoSum(nums []int, target int) []int {
    // 创建一个map来存储数字和对应的索引
    numMap := make(map[int]int)
    
    for i, num := range nums {
        // 计算需要的补数
        complement := target - num
        
        // 检查补数是否已经在map中
        if idx, exists := numMap[complement]; exists {
            // 如果存在，返回补数的索引和当前索引
            return []int{idx, i}
        }
        
        // 将当前数字和索引存入map
        numMap[num] = i
    }
    
    // 如果没有找到，返回空切片
    return nil
}

func main() {
    // 测试用例
    testCases := []struct {
        nums   []int
        target int
    }{
        {[]int{2, 7, 11, 15}, 9},
        {[]int{3, 2, 4}, 6},
        {[]int{3, 3}, 6},
        {[]int{1, 2, 3, 4, 5}, 9},
        {[]int{-1, -2, -3, -4, -5}, -8},
        {[]int{0, 4, 3, 0}, 0},
    }
    
    fmt.Println("两数之和测试:")
    fmt.Println("==================================================")
    
    for i, test := range testCases {
        result := twoSum(test.nums, test.target)
        fmt.Printf("测试用例 %d: nums=%v, target=%d -> 结果=%v\n", 
            i+1, test.nums, test.target, result)
    }
}
