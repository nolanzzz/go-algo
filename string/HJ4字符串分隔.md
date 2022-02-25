```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		segmentation(input.Text())
	}
}
func segmentation(s string) {
	for len(s) >= 8 {
		fmt.Println(s[:8])
		s = s[8:]
	}
	if len(s) == 0 {
		return
	}
	fmt.Print(s)
	for i := 0; i < 8-len(s); i++ {
		fmt.Print("0")
	}
	fmt.Println()
}
```