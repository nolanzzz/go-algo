```go
package main

import (
	"fmt"
)

func main() {
	var n float64
	fmt.Scan(&n)
	fmt.Println(int(n + 0.5)) // 或者math.Floor(n+0.5)
}

```