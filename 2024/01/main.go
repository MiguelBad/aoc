package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findDistance(left, right []int) int {
	var distance int

	for range left {
		leftNum := left[0]
		rightNum := right[0]

		left = left[1:]
		right = right[1:]

		d := leftNum - rightNum
		if d < 0 {
			d *= -1
		}

		distance += d
	}

	return distance
}

func quickSortPartition(arr []int, low, high int) int {
	idx := low - 1
	for i := low; i < high; i++ {
		if arr[i] < arr[high] {
			idx++
			arr[i], arr[idx] = arr[idx], arr[i]
		}
	}

	idx++
	arr[idx], arr[high] = arr[high], arr[idx]
	return idx
}

func quickSortHelper(arr []int, low, high int) {
	if low >= high {
		return
	}

	partition := quickSortPartition(arr, low, high)
	quickSortHelper(arr, low, partition-1)
	quickSortHelper(arr, partition+1, high)
}

func quickSort(arr []int) {
	quickSortHelper(arr, 0, len(arr)-1)
}

func readFile() ([]int, []int) {
	var (
		left  []int
		right []int
	)
	file, err := os.Open("input.txt")
	if err != nil {
		panic(fmt.Sprintf("failed to open file\n\nerr:\n%v\n", err))
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(strings.Trim(scanner.Text(), " "), " ")

		leftInt, err := strconv.Atoi(line[0])
		if err != nil {
			panic(fmt.Sprintf("failed to convert left string to int\n\nerr:\n%v\n", err))
		}
		rightInt, err := strconv.Atoi(line[len(line)-1])
		if err != nil {
			panic(fmt.Sprintf("failed to convert left string to int\n\nerr:\n%v\n", err))
		}

		left = append(left, leftInt)
		right = append(right, rightInt)
	}

	quickSort(left)
	quickSort(right)

	return left, right
}

func main() {
	left, right := readFile()
	distance := findDistance(left, right)
	fmt.Println(distance)
}
