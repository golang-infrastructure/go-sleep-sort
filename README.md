# 睡眠排序（Sleep Sort）

注意：此项目只是一个玩笑性质的项目，请勿将其用于生产环境！

# 一、安装

```bash
go get -u github.com/golang-infrastructure/go-sleep-sort
```

# 二、示例代码

```go
package main

import (
	"fmt"
	sleep_sort "github.com/golang-infrastructure/go-sleep-sort"
)

func main() {

	slice := []uint{8, 6, 9, 3}
	r := sleep_sort.SortUInt(slice)
	fmt.Println(r)
	// Output:
	// [3 6 8 9]

	slice = []uint{1000003, 1000010, 1000000}
	r = sleep_sort.SortUInt(slice)
	fmt.Println(r)
	// Output:
	// [1000000 1000003 1000010]

}
```





