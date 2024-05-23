# Recursion

当遇到树、链表、数组等递归结构的问题时，递归是一种常用的解决方法。
递归的本质是将一个大问题分解为一个或多个小问题，然后将小问题的解合并起来得到大问题的解。

## 需要用Recursion的地方
- Merge Sort
- Quick Sort
- Tree Traversal 
- Graph Traversal
- Dynamic Programming
- Backtracking
- Divide and Conquer


# Code
```go
package main

import "fmt"

func main() {
	res := ReverseString("hello")
	fmt.Println(res)
}

// 阶乘
func Factorial(value int) int {
	if value == 1 {
		return 1
	} else {
		return Factorial(value-1) * value
	}
}

// 斐波那契数列
func Fibonacci(value int) int {
	if value <= 1 {
		return value
	}
	return Fibonacci(value-1) + Fibonacci(value-2)
}

// 递归翻转字符串
func ReverseString(str string) string {
	if len(str) == 1 {
		return str
	} else {
		return ReverseString(str[1:]) + string(str[0])
	}
}
```

以 "hello" 为例，递归调用的展开和结果合并可以这样理解：

- 调用 `reverseString("hello")`
  - 调用 `reverseString("ello")`
    - 调用 `reverseString("llo")`
      - 调用 `reverseString("lo")`
        - 调用 `reverseString("o")`
          - 返回 `"o"`（基本情况）
        - 返回 `"o" + "l"` -> `"ol"`
      - 返回 `"ol" + "l"` -> `"oll"`
    - 返回 `"oll" + "e"` -> `"olle"`
  - 返回 `"olle" + "h"` -> `"olleh"`
