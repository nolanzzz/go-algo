方法一：KMP算法
```go
func strStr(haystack string, needle string) int {
    if len(needle) > len(haystack) {
        return -1
    }
    if len(needle) == 0 {
        return 0
    }
    // 使用kmp算法
    prefix := prefixTable([]byte(needle))
    m, n := len(haystack), len(needle)
    i, j := 0, 0
    for i < m {
        // match
        if j == n - 1 && needle[j] == haystack[i] {
            return i - j
        }
        if needle[j] == haystack[i] {
            i++; j++
        } else {
            j = prefix[j]
            if j == -1 {
                i++; j++
            }
        }
    }
    return -1
}

func prefixTable(needle []byte) []int {
    n := len(needle)
    prefix := make([]int, n)
    prefix[0] = 0
    i := 1
    l := 0
    for i < n {
        if needle[i] == needle[l] {
            l++
            prefix[i] = l
            i++
        } else {
            if l > 0 {
                l = prefix[l-1]
            } else {
                prefix[i] = 0
                i++
            }
        }
    }
    // 向右移位1
    for i := n - 1; i > 0; i-- {
        prefix[i] = prefix[i-1]
    }
    prefix[0] = -1
    return prefix
}
```

方法二：暴力解法-巧用切片
```go
func strStr(haystack string, needle string) int {
    if needle == "" {
        return 0
    }
    for i := 0; i <= len(haystack)-len(needle); i++ {
        if haystack[i] == needle[0] {
            if haystack[i:i+len(needle)] == needle {
                return i
            }
        }
    }
    return -1
}
```
