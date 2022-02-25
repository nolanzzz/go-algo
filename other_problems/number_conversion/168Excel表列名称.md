```go
func convertToTitle(columnNumber int) string {
    var res []rune
    for columnNumber > 0 {
        columnNumber--
        res = append(res, rune(columnNumber%26)+'A')
        columnNumber /= 26
    }
    return reverse(res)
}

func reverse(chars []rune) string {
    for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
        chars[i], chars[j] = chars[j], chars[i]
    }
    return string(chars)
}
```