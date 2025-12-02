// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

// 有效字符串需满足：

// 1.左括号必须用相同类型的右括号闭合。
// 2.左括号必须以正确的顺序闭合。
// 3.每个右括号都有一个对应的相同类型的左括号。

// 示例 2：

// 输入：s = "()[]{}"

// 输出：true

// 示例 3：

// 输入：s = "(]"

// 输出：false


package main		
import (
	// "package"	
	"fmt"
)

// isValid 判断括号字符串是否有效
func isValid(s string) bool {
	stack := []rune{}
	mapping := map[rune]rune{			
	')': '(',
	'}': '{',
	']': '[',
	}	
	// 遍历字符串中的每个字符
	for _, char := range s {
		// 如果是右括号，检查栈顶元素是否匹配
		if open, exists := mapping[char]; exists {
			// 栈为空或栈顶元素不匹配
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false			
			}		
			stack = stack[:len(stack)-1]	
		} else {
			// 如果是左括号，入栈
			stack = append(stack, char)
		}
	}
	// 最后栈应为空，表示所有括号匹配
	return len(stack) == 0
}

func main() {
	testStr := "{[(shkw728)]}"
	fmt.Printf("isValid(%s) = %v\n", testStr, isValid(testStr))
}