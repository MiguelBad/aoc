package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	x int
	y int
}

type Dir struct {
	north *Pos
	east  *Pos
	south *Pos
	west  *Pos
}

func newDir() *Dir {
	return &Dir{
		north: &Pos{x: 0, y: -1},
		east:  &Pos{x: 1, y: 0},
		south: &Pos{x: 0, y: 1},
		west:  &Pos{x: -1, y: 0},
	}
}

func throwError(desc string, err error) {
	panic(fmt.Sprintf("failed to %v\n\nerr\n%v\n", desc, err))
}

func readFile() ([][]string, *Pos) {
	file, err := os.Open("input.txt")
	if err != nil {
		throwError("open file", err)
	}

	var (
		playerPos = &Pos{}
		grid      = [][]string{}
	)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Err() != nil {
			throwError("scan file", scanner.Err())
		}

		var row []string
		for idx, char := range scanner.Text() {
			if string(char) == "^" {
				playerPos.x = idx
				playerPos.y = len(grid)
				row = append(row, ".")
				continue
			}
			row = append(row, string(char))
		}
		grid = append(grid, row)
	}

	return grid, playerPos
}

func isInBounds(grid *[][]string, pos *Pos) bool {
	if pos.x > len((*grid)[0])-1 || pos.x < 0 ||
		pos.y > len(*grid)-1 || pos.y < 0 {
		return false
	}
	return true
}

func isInObstacle(grid *[][]string, pos *Pos) bool {
	// fmt.Println(pos)
	if (*grid)[pos.y][pos.x] == "#" {
		return true
	}
	return false
}

func nextPos(currDir string, currPos *Pos, dir *Dir) *Pos {
	next := &Pos{x: currPos.x, y: currPos.y}
	switch currDir {
	case "n":
		next.y += dir.north.y
	case "e":
		next.x += dir.east.x
	case "s":
		next.y += dir.south.y
	case "w":
		next.x += dir.west.x
	}
	return next
}

func posCounter(posCount *int, grid *[][]string, playerPos *Pos, dir *Dir, visited [][]bool) {
	currDir := "n"
	for {
		next := nextPos(currDir, playerPos, dir)
		if !isInBounds(grid, next) {
			break
		}

		if isInObstacle(grid, next) {
			switch currDir {
			case "n":
				currDir = "e"
			case "e":
				currDir = "s"
			case "s":
				currDir = "w"
			case "w":
				currDir = "n"
			}

			next = nextPos(currDir, playerPos, dir)
			if !isInBounds(grid, next) {
				break
			}
		}

		if !visited[playerPos.y][playerPos.x] {
			visited[playerPos.y][playerPos.x] = true
			*posCount++
		}
		playerPos = next
	}
}

func main() {
	var (
		grid, playerPos = readFile()
		posCount        = 1
		dir             = newDir()
		visited         = make([][]bool, len(grid))
	)
	for i := range len(grid) {
		visited[i] = make([]bool, len(grid[0]))
	}

	posCounter(&posCount, &grid, playerPos, dir, visited)
	fmt.Println(posCount)
}
