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

func part1helper(rule *map[int][]int, update *[]int, sum *int) {
	for i := 0; i < len(*update)-1; i++ {
		if _, ok := (*rule)[(*update)[i]]; !ok {
			return
		}
		for j := i + 1; j < len(*update)-1; j++ {
			if !contains((*rule)[(*update)[i]], (*update)[j]) {
				return
			}
		}
	}
	*sum += (*update)[len(*update)/2]
}

func part1(rule *map[int][]int, updates *[][]int) int {
	var sum int
	for i := 0; i < len(*updates); i++ {
		part1helper(rule, &(*updates)[i], &sum)
	}
	return sum
}

func main() {
	rule, updates := readFile()
	fmt.Println(part1(&rule, &updates))
}
