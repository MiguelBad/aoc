package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part2Helper(report []int) bool {
	if part1Helper(report) {
		return true
	}

	for i := 0; i < len(report); i++ {
		modifiedReport := append([]int{}, report[:i]...)
		modifiedReport = append(modifiedReport, report[i+1:]...)
		if part1Helper(modifiedReport) {
			return true
		}
	}

	return false
}

func part2(reports [][]int) int {
	var safe int
	for _, report := range reports {
		if part2Helper(report) {
			fmt.Println(report)
			safe++
		}
	}
	return safe
}

func part1Helper(report []int) bool {
	isAscending := report[0] < report[1]
	for i := 1; i < len(report)-1; i++ {
		if isAscending {
			if report[i] > report[i+1] {
				return false
			}
		} else {
			if report[i] < report[i+1] {
				return false
			}
		}
		prevDiff := report[i-1] - report[i]
		nextDiff := report[i] - report[i+1]
		if prevDiff < 0 {
			prevDiff *= -1
		}
		if nextDiff < 0 {
			nextDiff *= -1
		}

		if prevDiff < 1 || prevDiff > 3 || nextDiff < 1 || nextDiff > 3 {
			return false
		}
	}
	return true
}

func part1(reports [][]int) int {
	var safe int
	for _, report := range reports {
		if part1Helper(report) {
			safe++
		}
	}
	return safe
}

func readFile() [][]int {
	var reports [][]int

	file, err := os.Open("input.txt")
	if err != nil {
		panic(fmt.Sprintf("failed to open file\n\nerr:\n%v\n", err))
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var report []int
		for _, i := range strings.Split(scanner.Text(), " ") {
			num, err := strconv.Atoi(i)
			if err != nil {
				panic(fmt.Sprintf("fialed to convert string to int\n\nerr:\n%v\n", err))
			}

			report = append(report, num)
		}

		reports = append(reports, report)
	}

	return reports
}

func main() {
	reports := readFile()
	safe := part1(reports)
	fmt.Println(safe)

	foo := part2(reports)
	fmt.Println(foo)
}
