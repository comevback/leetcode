// コーディングスキルチェック 地価の予想
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	Solution_Array()
	// Solution_LinkedList()
}

// define the listnode of distance list
type ListNode struct {
	Val   float64
	Price int
	Next  *ListNode
}

// ================================== use Array to store distance List ======================================
func Solution_Array() {
	var x, y int
	var k int
	var N int
	fmt.Scanf("%d %d", &x, &y)
	fmt.Scan(&k)
	fmt.Scan(&N)

	var distanceList []float64
	var hsMap map[float64]int = make(map[float64]int)
	var currentDis float64

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		arr := strings.Fields(text)

		x1, _ := strconv.Atoi(arr[0])
		y1, _ := strconv.Atoi(arr[1])
		price, _ := strconv.Atoi(arr[2])

		currentDis = math.Sqrt(float64(((x-x1)*(x-x1) + (y-y1)*(y-y1))))
		hsMap[currentDis] = price

		// sort the arr distanceList when insert, O(N)
		// if not, can sort it using sort.Float64s(distanceList) after the scanner loop
		// but the O() will be O(NlogN)
		inserted := false
		for i := 0; i < len(distanceList); i++ {
			if currentDis < distanceList[i] {
				// temp := distanceList[i:]  // Wrong, will change with the change of the distanceList
				// ***********************
				temp := make([]float64, len(distanceList[i:]))
				copy(temp, distanceList[i:])
				// ***********************
				distanceList = append(distanceList[:i], currentDis)
				distanceList = append(distanceList, temp...)
				inserted = true
				break
			}
		}
		if !inserted {
			distanceList = append(distanceList, currentDis)
		}

		N -= 1
		if N < 1 {
			break
		}
	}

	// sort.Float64s(distanceList)
	var sum float64 = 0
	for i := 0; i < k; i++ {
		sum += float64(hsMap[distanceList[i]])
	}

	avr := math.Round(sum / float64(k))
	fmt.Println(avr)
}

// ================================= use LinkedList to store distance List ===============================

func Solution_LinkedList() {
	var x, y int
	var k int
	var N int
	fmt.Scanf("%d %d", &x, &y)
	fmt.Scan(&k)
	fmt.Scan(&N)

	var distanceList *ListNode
	var currentDis float64

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		arr := strings.Fields(text)

		x1, _ := strconv.Atoi(arr[0])
		y1, _ := strconv.Atoi(arr[1])
		price, _ := strconv.Atoi(arr[2])

		currentDis = math.Sqrt(float64(((x-x1)*(x-x1) + (y-y1)*(y-y1))))
		newNode := &ListNode{
			Val:   currentDis,
			Price: price,
		}

		// if distanceList is empty
		if distanceList == nil {
			distanceList = newNode
		} else {
			inserted := false
			current := distanceList
			var parentNode *ListNode
			for current != nil {
				// currentDis bigger than this node, pass
				if currentDis > current.Val {
					parentNode = current
					current = current.Next
				} else {
					// currentDis smaller than this node
					newNode.Next = current
					// if this is not head, just insert
					if parentNode != nil {
						parentNode.Next = newNode
					} else {
						// if this is head, update the distanceList head to newNode
						distanceList = newNode
					}
					inserted = true
					break
				}
			}
			// currentDis bigger than all the existing node
			if !inserted {
				parentNode.Next = newNode
			}
		}

		N -= 1
		if N < 1 {
			break
		}
	}

	sum := 0
	kTemp := k
	current := distanceList
	for kTemp > 0 {
		sum += current.Price
		current = current.Next
		kTemp -= 1
	}

	avr := math.Round(float64(sum) / float64(k))
	fmt.Println(avr)
}
