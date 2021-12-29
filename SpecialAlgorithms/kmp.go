package main

import "fmt"

func main() {
	pattern := []byte("ABABC")
	text := []byte("ABABABABCAIABABCSS")
	// 1. 生成pattern的前缀表
	prefix := prefixTable(pattern)
	fmt.Println(prefix)
	// 2. 匹配
	kmpSearch(pattern, text)
}

func kmpSearch(pattern, text []byte) {
	prefix := prefixTable(pattern)
	m, n := len(pattern), len(text)
	i, j := 0, 0
	for j < n {
		// 找到match: i已经遍历到最后一位并且相等
		if i == m-1 && pattern[i] == text[j] {
			fmt.Println("Found a match at: ", j-i)
			i = prefix[i]
		}
		if pattern[i] == text[j] {
			i++
			j++
		} else {
			i = prefix[i]
			if i == -1 {
				i++
				j++
			}
		}
	}
}

func prefixTable(pattern []byte) []int {
	n := len(pattern)
	prefix := make([]int, n)
	prefix[0] = 0
	l := 0
	i := 1
	for i < n {
		if pattern[i] == pattern[l] {
			l++
			prefix[i] = l
			i++
		} else {
			if l == 0 {
				prefix[i] = 0
				i++
			} else {
				l = prefix[l-1]
			}
		}
	}
	// 将前缀表往后移一位，方便搜索时的移位操作
	for i := n - 1; i > 0; i-- {
		prefix[i] = prefix[i-1]
	}
	prefix[0] = -1
	return prefix
}
