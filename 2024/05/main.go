package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile() (map[int][]int, [][]int) {
	var updates [][]int
	rule := make(map[int][]int)

	file, err := os.Open("input.txt")
	if err != nil {
		panic(fmt.Sprintln("failed to open input", err))
	}
	defer file.Close()

	var updatePart bool
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Err() != nil {
			panic(fmt.Sprintln("failed to scan", scanner.Err()))
		}

		if scanner.Text() == "" {
			updatePart = true
			continue
		}

		if !updatePart {
			var val []int
			for _, i := range strings.Split(scanner.Text(), "|") {
				if v, err := strconv.Atoi(i); err == nil {
					val = append(val, v)
				}
			}
			if _, ok := rule[val[0]]; !ok {
				rule[val[0]] = []int{val[1]}
			} else {
				rule[val[0]] = append(rule[val[0]], val[1])
			}
		} else {
			var update []int
			for _, i := range strings.Split(scanner.Text(), ",") {
				if v, err := strconv.Atoi(i); err == nil {
					update = append(update, v)
				}
			}
			updates = append(updates, update)
		}
	}

	return rule, updates
}

func contains(arr []int, item int) bool {
	for i := 0; i < len(arr); i++ {
		if (arr)[i] == item {
			return true
		}
	}
	return false
}

func reOrder(rule *map[int][]int, arr *[]int) *[]int {
	var ordered []int
	count := make(map[int]int)
	for i := 0; i < len(*arr); i++ {
		for j := 0; j < len(*arr); j++ {
			if i == j {
				continue
			}

			if contains((*rule)[(*arr)[i]], (*arr)[j]) {
				if _, ok := count[(*arr)[i]]; !ok {
					count[(*arr)[i]] = 1
				} else {
					count[(*arr)[i]]++
				}
			}
		}
	}

	for len(count) > 0 {
		// largest[0] = count
		// largest[1] = key
		largest := []int{0, 0}
		for i := range count {
			if count[i] > largest[0] {
				largest[0] = count[i]
				largest[1] = i
			}
		}
		ordered = append(ordered, largest[1])
		delete(count, largest[1])
	}
	return &ordered
}

func checkOrderHelper(rule *map[int][]int, update *[]int, correctSum *int, incorrectSum *int) {
	var incorrect bool
	for i := 0; i < len(*update)-1; i++ {
		if incorrect {
			break
		}

		if _, ok := (*rule)[(*update)[i]]; !ok {
			incorrect = true
			break
		}

		for j := i + 1; j < len(*update)-1; j++ {
			if !contains((*rule)[(*update)[i]], (*update)[j]) {
				incorrect = true
				break
			}
		}
	}

	if incorrect {
		ordered := *reOrder(rule, update)
		*incorrectSum += ordered[len(ordered)/2]
	} else {
		*correctSum += (*update)[len(*update)/2]
	}
}

func checkOrder(rule *map[int][]int, updates *[][]int) (int, int) {
	var correctSum int
	var incorrectSum int
	for i := 0; i < len(*updates); i++ {
		checkOrderHelper(rule, &(*updates)[i], &correctSum, &incorrectSum)
	}
	return correctSum, incorrectSum
}

func main() {
	rule, updates := readFile()
	correct, incorrect := checkOrder(&rule, &updates)
	fmt.Println("Correct sum:", correct)
	fmt.Println("Incorrect sum:", incorrect)
}
