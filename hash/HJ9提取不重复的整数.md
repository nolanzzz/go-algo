```go
package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    m := make(map[int]bool)
    for n > 0 {
        if m[n%10] {
            n /= 10
            continue
        }
        fmt.Print(n % 10)
        m[n % 10] = true
        n /= 10
    }
}
```