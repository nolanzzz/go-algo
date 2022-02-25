```go
func titleToNumber(columnTitle string) int {
    // 26进制转10进制
    var res int
    for _, char := range columnTitle {
        res = res * 26 + int(char-'A'+1)
    }
    return res
}

```