方法一：手动换算
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
        fmt.Println(conversion(input.Text()))
    }
}

func conversion(s string) int {
    var res int
    digits := s[2:]
    dict := map[rune]int{'0':0,'1':1,'2':2,'3':3,'4':4,'5':5,'6':6,'7':7,'8':8,'9':9,'A':10,'B':11,'C':12,'D':13,'E':14,'F':15}
    for _, digit := range digits {
        res = res * 16 + dict[digit]
    }
    return res
}
```

方法二：直接用scan方法获取10进制数
```go
package main

import "fmt"

func main() {
	var num int
	for {
		_, err := fmt.Scanf("0x%X", &num)
		if err != nil {
			break
		}
		fmt.Println(num)
	}
}
```