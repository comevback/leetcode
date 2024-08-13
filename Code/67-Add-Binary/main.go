package main

func main() {

}

func addBinary(a string, b string) string {
	res := "" // 存储最终的结果字符串

	aIndex, bIndex := len(a)-1, len(b)-1 // 从字符串的末尾开始处理
	add := false                         // 用于跟踪进位（是否需要在下一位上加 1）

	for aIndex >= 0 || bIndex >= 0 || add { // 当 a 或 b 还有位没有处理，或仍然有进位时继续循环
		thisNum := 0 // thisNum 用于保存当前位的计算结果

		// 如果 a 还有剩余的位，处理 a 的当前位
		if aIndex >= 0 {
			if a[aIndex] == '1' { // 如果 a 的当前位是 '1'
				thisNum += 1 // 加上 1
			}
			aIndex -= 1 // 移动到 a 的前一位
		}

		// 如果 b 还有剩余的位，处理 b 的当前位
		if bIndex >= 0 {
			if b[bIndex] == '1' { // 如果 b 的当前位是 '1'
				thisNum += 1 // 加上 1
			}
			bIndex -= 1 // 移动到 b 的前一位
		}

		// 如果之前有进位，加 1
		if add {
			thisNum += 1
		}

		// 如果当前位的和大于或等于 2，需要进位
		if thisNum >= 2 {
			add = true   // 设置进位标志
			thisNum -= 2 // 结果中只保留当前位的 0 或 1
		} else {
			add = false // 没有进位
		}

		// 将当前位的结果（0 或 1）加到结果字符串的前面
		if thisNum == 1 {
			res = "1" + res
		} else if thisNum == 0 {
			res = "0" + res
		}
	}

	return res // 返回最终的结果字符串
}
