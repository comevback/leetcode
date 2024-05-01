package main

import (
	"fmt"
	"strconv"
)

func main() {
	var chars []byte = []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'c', 'c', 'c'}
	fmt.Println(compress(chars))
	fmt.Println(chars)
}

// ================================================================================================================================
// 1. 双指针 O(n)-O(n) 空间复杂度高，用了多余的数据结构
// 每当遇到新的字符时，开始技术，直到下次遇到另外的字符，将统计的当前字符和字符的数量写入结果字符串，然后当前字符变成新的字符，数量变成1，重新开始统计
// 当遍历结束时，将最后一个字符和数量写入结果字符串
// 最后将结果字符串转换成字符数组，复制到原数组中，返回字符串长度
func compress1(chars []byte) int {
	var fast int = 0
	var count int = 0
	var current byte
	var res string
	for fast < len(chars) {
		if chars[fast] != current {
			if count == 0 {
				current = chars[fast]
				count = 1
				fast += 1
				continue
			}
			if count == 1 {
				res += string(current)
				current = chars[fast]
				fast += 1
				continue
			}
			res += string(current)
			res += strconv.Itoa(count)
			count = 1
			current = chars[fast]
		} else {
			count += 1
		}
		fast += 1
	}
	res += string(current)
	if count != 1 {
		res += strconv.Itoa(count)
	}

	for i := 0; i < len(res); i++ {
		chars[i] = res[i]
	}

	return len(res)
}

// ================================================================================================================================
// 2. 双指针 从后往前遍历 错误 无法处理原数组的长度变化
func compress2(chars []byte) int {
	var head int = len(chars) - 1
	var tail int = len(chars) - 1
	var count int = 0
	var current byte = chars[tail]

	for head >= 0 {
		if chars[head] != current {
			if count == 1 {
				copy(chars[head+1:tail+1], []byte{current})
			} else {
				add := []byte{current}
				add = append(add, []byte(strconv.Itoa(count))...)
				copy(chars[head+1:tail+1], add)
				for i := head + len(add) + 1; i < len(chars); i++ {
					chars[i] = chars[i+count-len(add)]
				}
			}
			tail = head
			current = chars[head]
			count = 1
		} else {
			count += 1
		}
		head -= 1
	}

	if count == 1 {
		copy(chars[:tail+1], []byte{current})
	} else {
		add := []byte{current}
		add = append(add, []byte(strconv.Itoa(count))...)
		copy(chars[head+1:tail+1], add)
	}

	return len(chars)

}

// ================================================================================================================================
// 碰到要对数组进行in-place操作时，不要想着增删数组，或者copy之类的，而是要用双指针把结果元素从头堆进数组里，替换掉原数组
// 最后原数组的前n个元素，就是新数组
//
// 3.双指针 in-place O(n)-O(1)
func compress(chars []byte) int {

	// 边界条件：如果数组为空，直接返回0
	if len(chars) == 0 {
		return 0
	}

	// tail指针指向新数组的尾部，count记录当前字符的数量
	tail := 0
	count := 1

	// 遍历数组
	for i := 0; i < len(chars); i++ {

		// 如果当前字符和下一个字符不相等，或者已经遍历到最后一个字符
		if i == len(chars)-1 || chars[i] != chars[i+1] {
			// 将当前字符写入新数组
			chars[tail] = chars[i]
			// tail指针后移
			tail += 1
			// 如果当前字符数量大于1，将数量也写入新数组
			if count > 1 {
				// 将数字转换成字符串，然后转换成byte数组，把转换成的byte的字符数量count写入新数组
				numStr := strconv.Itoa(count)
				numByte := []byte(numStr)
				for _, value := range numByte {
					chars[tail] = value
					tail += 1
				}
			}
			// 重置count
			count = 1
		} else {
			// 如果当前字符和下一个字符相等，数量加1
			count += 1
		}
	}

	// 返回新数组的长度
	return tail
}

// ================================================================================================================================
// 4.answer from chatgpt
func compress_chatgpt(chars []byte) int {
	if len(chars) == 0 {
		return 0
	}

	tail, count := 0, 1
	// 正向遍历更自然地处理这种情况
	for i := 1; i <= len(chars); i++ {
		if i == len(chars) || chars[i] != chars[i-1] {
			chars[tail] = chars[i-1]
			tail++
			if count > 1 {
				// 将数字转为字符串后转为字节切片，并复制到chars中
				for _, c := range strconv.Itoa(count) {
					chars[tail] = byte(c)
					tail++
				}
			}
			count = 1 // 重置计数
		} else {
			count++
		}
	}
	return tail // tail现在是新的数组长度
}

// ================================================================================================================================
// 5.answer from others in leetcode
func cp(chars []byte) int {
	// Initialize pointers and counters
	writeIndex := 0
	readIndex := 0

	for readIndex < len(chars) {
		// Initialize the current character and its count
		currentChar := chars[readIndex]
		count := 0

		// Count consecutive repeating characters
		for readIndex < len(chars) && chars[readIndex] == currentChar {
			readIndex++
			count++
		}

		// Write the compressed character
		chars[writeIndex] = currentChar
		writeIndex++

		// If count is greater than 1, write its digits
		if count > 1 {
			countStr := fmt.Sprintf("%d", count)
			for _, digit := range countStr {
				chars[writeIndex] = byte(digit)
				writeIndex++
			}
		}
	}

	return writeIndex
}

func compress_leetcode(chars []byte) int {
	compressedLength := cp(chars)
	fmt.Println("Compressed array:", string(chars[:compressedLength])) // Output: a2b2c3
	return compressedLength
}
