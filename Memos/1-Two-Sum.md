# 1. Two Sum

Easy

Hint
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

 

Example 1:
> Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Example 2:
> Input: nums = [3,2,4], target = 6
Output: [1,2]

Example 3:
> Input: nums = [3,3], target = 6
Output: [0,1]
 

Constraints:
> 2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
Only one valid answer exists.
 
Follow-up: Can you come up with an algorithm that is less than O(n2) time complexity?

---

## My First Try (√) 双重遍历
```go
// 第一种方法，两层遍历
func twoSum(nums []int, target int) []int {
	for i, v1 := range nums {
		// 第二层，记得从第一层的下一个元素开始，
		for j := i + 1; j < len(nums); j++ {
			if v1+nums[j] == target && i != j {
				return []int{i, j}
			}
		}
	}
	// 如果没有结果，返回一个空的切片
	return []int{}
}
```

## 用元素值作为Map的key，索引作为value，找 map[target - val] 

```go
// 第二种方法，用哈希表map,用元素值作为Map的key，索引作为value，找map[target - val]是否存在
func twoSum2(nums []int, target int) []int {

	// 定义一个map，key的值是元素值，value的值是切片的索引
	var hsMap map[int]int = make(map[int]int)

	// 遍历所有元素，如果在当前循环中，找到了以val的互补数(target-val)为key的map元素，也就是这个值在map中存在，那么返回这两个元素的下标
	for i, val := range nums {

		// 找当前值的互补值，是否在map中，如果在，取其下标和当前遍历的下标i组合，返回
		if j, exist := hsMap[target-val]; exist {
			return []int{i, j}
		}
		//如果没有，把当前值val作为key，i作为value，储存到map中
		hsMap[val] = i
	}

	// 如果没有结果，返回空切片
	return []int{}
}
```