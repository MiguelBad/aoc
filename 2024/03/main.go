package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getProductSum(multiplicand [][]int) int {
	var productSum int
	for _, m := range multiplicand {
		productSum += m[0] * m[1]
	}
	return productSum
}

func getMultiplicand(inst string) [][]int {
	var cleaned [][]int
	var currDigit string
	var multiplicand []int
	openningPattern := "mul("

	var disabled bool

	fmt.Println(string([]byte{inst[0], inst[1]}))

	for i, v := range inst {
		if string(v) == "d" {
			if i > len(inst)-1 {
				continue
			}

			if string([]byte{inst[i], inst[i+1], inst[i+2], inst[i+3]}) == "do()" {
                disabled = false
			} else if string([]byte{inst[i], inst[i+1], inst[i+2], inst[i+3], inst[i+4], inst[i+5], inst[i+6]}) == "don't()" {
                disabled = true
			}
		}

		if disabled {
			continue
		}

		if len(openningPattern) > 0 {
			if rune(openningPattern[0]) == v {
				openningPattern = openningPattern[1:]
				continue
			} else {
				openningPattern = "mul("
				continue
			}
		} else {
			if _, err := strconv.Atoi(string(v)); err == nil {
				if len(currDigit) < 3 {
					currDigit += string(v)
					continue
				}
				currDigit = ""
				openningPattern = "mul("
				continue
			}

			if string(v) == "," && isValidDig(currDigit) {
				multiplicand = append(multiplicand, toInt(currDigit))
				currDigit = ""
				continue
			}

			if string(v) == ")" && isValidDig(currDigit) {
				multiplicand = append(multiplicand, toInt(currDigit))
				cleaned = append(cleaned, multiplicand)
			}

			openningPattern = "mul("
			currDigit = ""
			multiplicand = []int{}
		}
	}

	return cleaned
}

func isValidDig(s string) bool {
	if v, err := strconv.Atoi(s); err == nil {
		return v > 0 && v < 1000
	} else {
		return false
	}
}

func toInt(s string) int {
	if v, err := strconv.Atoi(s); err == nil {
		return v
	} else {
		panic("failed to convert")
	}
}

func readFile() string {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(fmt.Sprintf("failed to open file\n\nerr:\n%v\n", err))
	}

	return strings.Trim(string(file), "\n")
}

func main() {
	inst := readFile()
	multiplicand := getMultiplicand(inst)
	fmt.Println(getProductSum(multiplicand))
}
