package main

import (
	"strconv"
)

func main() {
}

func Test1(strNum string) (int, error) {
	if _, err := strconv.Atoi(strNum); err != nil {
		return 0, err
	}
	result := 0
	for i := 0; i < len(strNum); i++ {
		num, _ := strconv.Atoi(string(strNum[i]))
		result = result + num
	}
	strResult := strconv.Itoa(result)
	for {
		if len(strResult) == 1 {
			break
		}
		result = 0
		for i := 0; i < len(strResult); i++ {
			num, _ := strconv.Atoi(string(strResult[i]))
			result = result + num
		}
		strResult = strconv.Itoa(result)
	}
	return result, nil
}
