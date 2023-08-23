package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	result, err := UnpackingFunc("a4b5r8\\9\\56\\\\7")
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(result)
}

func UnpackingFunc(str string) (string, error) {
	var (
		result string
	)
	if str == "" {
		return "", nil
	}
	if str[0] < 97 || str[0] > 122 {
		return result, fmt.Errorf("incorrent string")
	}
	for i := 0; i <= len(str)-1; i++ {
		if str[i] >= 97 && str[i] <= 122 {
			result += string(str[i])
		} else if str[i] == '\\' {
			result += string(str[i+1])
			i++
		} else {
			arr := make([]byte, 0)
			for j := i; j <= len(str)-1; j++ {
				if str[j] >= 48 && str[j] <= 57 {
					arr = append(arr, str[j])
				} else {
					break
				}
			}
			atoi, err := strconv.Atoi(string(arr))
			if err != nil {
				return "", err
			}
			tmp := strings.Repeat(string(str[i-1]), atoi-1)
			result += tmp
		}
	}
	return result, nil
}
