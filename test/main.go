package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	exec()
}

// main function
func exec() {
	var input int
	_, err := fmt.Scan(&input)
	if err != nil {
		log.Fatal("invalid input: ", err)
	}
	res := callCal(input, []int{})
	printArr(res)
}

// calculate the sum of the square of the positive numbers
func Sum(arr []string) int {
	if len(arr) == 0 {
		return 0
	} else {
		num, err := strconv.Atoi(arr[0])
		if err != nil {
			log.Fatal(err)
		}
		if num <= 0 {
			return Sum(arr[1:])
		} else {
			return Square(num) + Sum(arr[1:])
		}
	}
}

func Square(n int) int {
	return n * n
}

// call the callSum function for certain times
func callCal(times int, results []int) []int {
	if times == 0 {
		return results
	} else {
		arr := getInput()
		if len(arr) > 1 || len(arr) == 0 {
			log.Fatal("invalid number")
		}
		num, err := strconv.Atoi(arr[0])
		if err != nil {
			log.Fatal(err)
		}
		// defer fmt.Println(callSum(num))
		results = append(results, callSum(num))
		return callCal(times-1, results)
	}
}

// calculate the sum of given numbers
func callSum(amount int) int {
	arr := getInput()
	if len(arr) != amount {
		log.Fatal("not right nums")
	}
	return Sum(arr)
}

// getInput get input from stdin for the callSum and callCal
func getInput() []string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	str = strings.TrimSpace(str)
	arr := strings.Split(str, " ")
	return arr
}

// print the result in thn end
func printArr(nums []int) {
	if len(nums) != 0 {
		fmt.Println(nums[0])
		printArr(nums[1:])
	}
}
