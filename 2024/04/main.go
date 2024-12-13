package main

import (
	"bufio"
	"fmt"
	"os"
)

func getDirList() map[string][]int {
	var dirList map[string][]int = make(map[string][]int)
	dirList["nw"] = []int{-1, -1}
	dirList["n"] = []int{0, -1}
	dirList["ne"] = []int{1, -1}
	dirList["e"] = []int{1, 0}
	dirList["se"] = []int{1, 1}
	dirList["s"] = []int{0, 1}
	dirList["sw"] = []int{-1, 1}
	dirList["w"] = []int{-1, 0}
	return dirList
}

func getNext(curr, next []int, textList *[][]string) ([]int, bool) {
	var isErr bool
	x := curr[0] + next[0]
	if x > len((*textList)[0])-1 || x < 0 {
		isErr = true
	}
	y := curr[1] + next[1]
	if y > len(*textList)-1 || y < 0 {
		isErr = true
	}
	return []int{x, y}, isErr
}

func searchHelper(dirList map[string][]int, dir string, curr []int, textList *[][]string, word string, count *int) {
	fmt.Println(curr[0], curr[1], word, (*textList)[curr[0]][curr[1]], dir)
	if word == "" {
		*count++
		return
	}

	if dir != "" {
		nextPos, err := getNext(curr, dirList[dir], textList)
		if err {
			return
		}
		if (*textList)[nextPos[1]][nextPos[0]] == string(word[0]) {
			searchHelper(dirList, dir, nextPos, textList, word[1:], count)
		}
		return
	}

	for key, val := range dirList {
		nextPos, err := getNext(curr, val, textList)
		if err {
			continue
		}
		if (*textList)[nextPos[1]][nextPos[0]] == string(word[0]) {
			searchHelper(dirList, key, nextPos, textList, word[1:], count)
		}
	}
}

func search(textList *[][]string, dirList map[string][]int) int {
	var count int
	for rIdx, row := range *textList {
		for lIdx, letter := range row {
			if letter == "X" {
				curr := []int{lIdx, rIdx}
				fmt.Println(curr, "X")
				searchHelper(dirList, "", curr, textList, "MAS", &count)
				fmt.Println()
			}
		}
	}
	return count
}

func readFile() [][]string {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(fmt.Sprintf("failed to read file\n\nerr:\n%v\n", err))
	}
	defer file.Close()

	var textList [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			panic(fmt.Sprintf("failed to scan line\n\nerr:\n%v\n", err))
		}

		column := []string{}
		for _, s := range scanner.Text() {
			column = append(column, string(s))
		}
		textList = append(textList, column)
	}

	return textList
}

func main() {
	dirList := getDirList()
	textList := readFile()
	fmt.Println(search(&textList, dirList))
}
