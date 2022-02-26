```go
package main

import (
    "fmt"
    "sort"
)

func main() {
    var n, k, v int
    fmt.Scan(&n)
    m := make(map[int]int)
    var keys []int
    for n > 0 {
        fmt.Scanf("%d %d", &k, &v)
        m[k] += v
        n--
    }
    for k := range m {
        keys = append(keys, k)
    }
    sort.Ints(keys)
    for _, key := range keys {
        fmt.Printf("%d %d\n", key, m[key])
    }
}
```