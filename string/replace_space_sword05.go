package str

import (
	"strings"
	"unicode"
)

//剑指 Offer 05. 替换空格
//请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
//
//示例 1：
//输入：s = "We are happy."
//输出："We%20are%20happy."

func replaceSpace(s string) string {
	var count int
	// 获取空格的个数
	for _, v := range s {
		if unicode.IsSpace(v) {
			count++
		}
	}

	str := &strings.Builder{}
	// 预分配内存, 如果是字符串长度可变，在原s基础上修改s，则可以倒序遍历s，减少移动
	str.Grow(len(s) + 2*count)
	for i := range s {
		if s[i] == ' ' {
			str.Write([]byte("%20"))
			continue
		}
		str.WriteByte(s[i])
	}
	return str.String()
}
