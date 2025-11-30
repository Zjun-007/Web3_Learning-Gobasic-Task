//  最长公共前缀
// 编写一个函数来查找字符串数组中的最长公共前缀。

// 如果不存在公共前缀，返回空字符串 "".

// 提示：

// 1 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] 如果非空，则仅由小写英文字母组成


// 字符串  // 切片

// 函数
// 循环--得到相同的前缀

package main
import "fmt"

func LongestCommonPrefix(strs []string) string {
	// 边界条件
	if len(strs) == 0 {
		return ""
	}				
	prefix := strs[0]					

	for i := 1; i < len(strs); i++ {		
		// 循环得到相同的前缀
		// strs := []string{"flower", "flow", "flight"}   len(prefix) = 6
		// i = 1  strs[1] = "flow"   len(strs[1]) = 4
		// prefix = prefix[:len(prefix)]	  = "flower"  != strs[1][:4] = "flow"
		// prefix = prefix[:len(prefix)-1]  = "flowe"   != strs[1][:4] = "flow"
		// prefix = prefix[:len(prefix)-1]  = "flow"    == strs[1][:4] = "flow"	
		// prefix = "flow"
		// i = 2  strs[2] = "flight"   len(strs[2]) = 6
		// prefix = prefix[:len(prefix)]	  = "flow"    != strs[2][:4] = "flig"	//prefix=最近的赋值flow
		// prefix = prefix[:len(prefix)-1]  = "flo"     != strs[2][:3] = "fli"
		// prefix = prefix[:len(prefix)-1]  = "fl"      == strs[2][:2] = "fl"
		// prefix = "fl"
		for len(prefix) > 0 && strs[i][:min(len(prefix), len(strs[i]))] != prefix {
			prefix = prefix[:len(prefix)-1]
		}
		if len(prefix) == 0 {
			return ""
		}				

	}				

	return prefix
}	

func main() {
	strs := []string{"flower", "flow", "flight"}
	result := LongestCommonPrefix(strs)
	fmt.Println(result) // 输出: "fl"
}			
