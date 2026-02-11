package main

import (
	"strings"
)

func main() {
}

func getPassword(grid [3][3]string, command []string) (string, error) {
	currentX := 1
	currentY := 0
	result := ""
	for i := 0; i < len(command); i++ {
		if strings.Contains(command[i], "down") {
			currentY += 1
			if currentY > 2 {
				currentY = 2
			}
			if strings.Contains(command[i], "downT") {
				result = result + grid[currentY][currentX]
			}
		} else if strings.Contains(command[i], "up") {
			currentY -= 1
			if currentY < 0 {
				currentY = 0
			}
			if strings.Contains(command[i], "upT") {
				result = result + grid[currentY][currentX]
			}
		} else if strings.Contains(command[i], "left") {
			currentX -= 1
			if currentX < 0 {
				currentX = 0
			}
			if strings.Contains(command[i], "leftT") {
				result = result + grid[currentY][currentX]
			}
		} else if strings.Contains(command[i], "right") {
			currentX += 1
			if currentX > 2 {
				currentX = 2
			}
			if strings.Contains(command[i], "rightT") {
				result = result + grid[currentY][currentX]
			}
		}
	}

	return result, nil
}
