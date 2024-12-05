package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkReportHelper(report []int) bool {
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

        fmt.Println(prevDiff, nextDiff)
	}
	return true
}

func checkReport(reports [][]int) int {
	var safe int
	for _, report := range reports {
		if checkReportHelper(report) {
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
	safe := checkReport(reports)
	fmt.Println(safe)
}
