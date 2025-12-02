
// 基本值类型

// 考察：数组操作、进位处理

// 题目：给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一

// 给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。
// 这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。

// 将大整数加 1，并返回结果的数字数组。

// 示例: 

// 输入：digits = [4,3,2,1]
// 输出：[4,3,2,2]
// 解释：输入数组表示数字 4321。
// 加 1 后得到 4321 + 1 = 4322。
// 因此，结果应该是 [4,3,2,2]。

package main	
import "fmt"	

func plusOne(digits []int) []int {
    n := len(digits)        
    // 从最后一位开始加1
    for i := n - 1; i >= 0; i-- {
        // 如果当前位小于9，直接加1并返回结果
        if digits[i] < 9 {
            digits[i]++
            return digits
        }
        // 当前位是9，加1后变成0，循环会继续向前进位      
        digits[i] = 0       
    }
    // 如果所有位都是9，说明进位后长度增加1
    // return append([]int{1}, digits...) 也可以先初始化1，然后增加切片
    newDigits := make([]int, n+1)
    newDigits[0] = 1
    return newDigits
}

func main() {
    testCases := [][]int{
        {4, 3, 2, 1},   
        {9, 9, 9},     
        {4, 3, 2, 9},       
        {0},             
        {1, 9, 9},      
    }
    for _, digits := range testCases {
        result := plusOne(digits)
        fmt.Printf("输入: %v -> 输出: %v\n", digits, result)
    }
}