// 136. 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。
// 找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
// 例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。

// 回文数

// 考察：数字操作、条件判断
// 题目：判断一个整数是否是回文数 
package main	

import (
    "fmt"
)

// 直接判断，区分大小写
func isPalindrome(s string) bool {
    left, right := 0, len(s)-1
    
    for left < right {
        if s[left] != s[right] {
            return false
        }
        left++
        right--
    }
    
    return true
}

func main() {
    testCases := []string{
        "racecar",           // true
        "hello",             // false
        "A man, a plan, a canal: Panama",  // true
        "Was it a car or a cat I saw?",     // true
        "No 'x' in Nixon",                  // true
        "12321",             // true
        "123ab321",          // false
        "",                  // true (空字符串是回文)
        "a",                 // true (单个字符是回文)
        "AbBa",              // true
    }
    
    fmt.Println("区分大小写，包括所有字符:")
    for _, test := range testCases {
        fmt.Printf("%-35q -> %v\n", test, isPalindrome(test))
    }
    
}