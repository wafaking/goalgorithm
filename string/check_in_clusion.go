package str

import (
	"fmt"
	"log"
)

//567. 字符串的排列
//给定两个字符串s1和s2，写一个函数来判断s2是否包含s1的排列。
//换句话说，第一个字符串的排列之一是第二个字符串的子串。

//示例 1：
//输入: s1 = "ab" s2 = "eidbaooo"
//输出: True
//解释: s2 包含 s1 的排列之一 ("ba").

//示例 2：
//输入: s1= "ab" s2 = "eidboaoo"
//输出: False

func checkInclusion(s1 string, s2 string) bool {
	// 1. 校验
	m, n := len(s1), len(s2)
	if m > n {
		return false
	}

	// 2. 初始化两个数组，以字母ascii码值为索引，出现的次数为count
	array1 := [26]int{}
	array2 := [26]int{}

	//"ab" "eidboaoo"
	for k, v := range s1 {
		array1[v-'a']++
		v2 := s2[k] // s2也只取两位
		array2[v2-'a']++
	}

	// 先比较一次，如果不相等，使用滑动窗口向右滑动
	if array2 == array1 {
		return true
	}

	//"ab" "eidboaoo"
	// 使用滑动窗口，逐渐向右滑动s2，
	//右侧新进入的字母次数+1,左侧滑出的字母次数-1
	for i := m; i < n; i++ { // 从未比较过的地方开始,如d开始（即窗口右侧）
		v2 := s2[i] - 'a' // 即d所对应的ascii值
		array2[v2]++      // d的索引的次数+1

		v2Pre := s2[i-m] - 'a' // 窗口左侧的e对应的ascii码值
		array2[v2Pre]--        // e索引对应的次数-1（窗口左侧）
		if array1 == array2 {
			return true
		}
	}
	return false
}

func checkInclusion1(s1, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}
	//"ab", "eidbaooo"
	var cnt1, cnt2 [26]int
	for i, ch := range s1 {
		fmt.Println("---", i, string(ch), ch-'a', "----", string(s2[i]), s2[i]-'a')
		cnt1[ch-'a']++
		cnt2[s2[i]-'a']++
	}

	log.Println("a1: ", cnt1, n)
	log.Println("a2: ", cnt2, m)

	if cnt1 == cnt2 {
		return true
	}

	fmt.Println("cnt1-------: ", cnt1)
	//"ab", "eidbaooo"
	for i := n; i < m; i++ { // i=n=2, m=8

		cnt2[s2[i]-'a']++
		fmt.Println("====1=: ", i, n, string(s2[i]), s2[i]-'a', cnt2)
		//fmt.Println("====2=: ", cnt2)

		cnt2[s2[i-n]-'a']--
		fmt.Println("====2=: ", cnt2)
		if cnt1 == cnt2 {
			return true
		}
	}
	return false
}

//这道题是 76. 最小覆盖子串 的简单版本。我们为大家罗列一些要点：

//排列不讲究顺序，但是字符出现的 种类 和 次数 要恰好对应相等；
//暴力解法做了很多重复和没有必要的工作；

//思考可以使用双指针（一前一后、交替向右边移动）的原因：少考虑很多暴力解法须要考虑的子区间，
//已经得到了一部分子区间的信息以后，很多子区间都不必要考虑，
//因此根据特定的情况右指针、左指针交替向右移动，从时间复杂度 O(N2)O(N^2)O(N2) 降到 O(N)O(N)O(N)；

//由于我们只关心子区间里的元素的种数和各个字符出现的次数。因此须要统计字符串 s1 出现的字符的种数和次数，
//和在字符串 s2 上的两个变量所确定的滑动窗口中出现的字符种数和次数；

//还须要设计一个变量 winCount，表示滑动窗口在 s2 上滑动的时候，出现在 s1 中的字符的种类数，规则如下：
//如果某一个字符出现的次数恰好等于 s1 中对应字符出现的次数，winCount += 1；
//在左边界向右移动的过程当中，某一个字符对应的次数减少以后，恰好小于了 s1 对应的字符出现的次数，winCount -= 1；
//当滑动窗口中出现的字符种类数和 s1 中出现的字符种类数相等（根据 winCount 的定义，对应次数也相等），
//并且 s2 上的滑动窗口的长度等于字符串 s1 的长度的时候，就找到了 s1 的一个排列。

//	winCount: 滑动窗口中的字符要满足字符出现的次数大于等于s1中对应字符出现的次数的时候才加1，
//	 不仅统计了种类，还代表了次数。使得我们可以通过 winCount 的数值去了解整个滑动窗口的信息。
