# 190. Reverse Bits

Solved
Easy
Topics
Companies
Reverse bits of a given 32 bits unsigned integer.

Note:

Note that in some languages, such as Java, there is no unsigned integer type. In this case, both input and output will be given as a signed integer type. They should not affect your implementation, as the integer's internal binary representation is the same, whether it is signed or unsigned.
In Java, the compiler represents the signed integers using 2's complement notation. Therefore, in Example 2 above, the input represents the signed integer -3 and the output represents the signed integer -1073741825.

Example 1:

Input: n = 00000010100101000001111010011100
Output: 964176192 (00111001011110000010100101000000)
Explanation: The input binary string 00000010100101000001111010011100 represents the unsigned integer 43261596, so return 964176192 which its binary representation is 00111001011110000010100101000000.
Example 2:

Input: n = 11111111111111111111111111111101
Output: 3221225471 (10111111111111111111111111111111)
Explanation: The input binary string 11111111111111111111111111111101 represents the unsigned integer 4294967293, so return 3221225471 which its binary representation is 10111111111111111111111111111111.

Constraints:

The input must be a binary string of length 32

Follow up: If this function is called many times, how would you optimize it?

---

# Code

```go
package main

func main() {

}

// reverseBits 反转一个 int32 类型整数的所有位
func reverseBits(num int32) int32 {
	var result int32 = 0
	for i := 0; i < 32; i++ {
		// 将 result 向左移一位，为下一位腾出空间
		result <<= 1
		// 获取 num 的最低位
		lowest := num & 1
		// 将这个位添加到 result 的最低位
		result |= lowest
		// 将 num 向右移动一位，准备下一次迭代
		num >>= 1
	}
	return result
}
```
