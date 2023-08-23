package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	dict := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	fmt.Println(dict)
	fmt.Println(Anagram(dict))
}

func uniqueInLower(str []string) []string {
	result := make([]string, 0, len(str))
	unique := make(map[string]bool)

	for _, i := range str {
		if !unique[i] {
			result = append(result, strings.ToLower(i))
			unique[i] = true
		}
	}
	return result
}

func Anagram(str []string) map[string][]string {
	if len(str) < 2 {
		return nil
	}
	buffer := make(map[string][]string)
	uniqIn := uniqueInLower(str)
	for _, i := range uniqIn {
		sorted := []rune(i)
		sort.Slice(sorted, func(i, j int) bool {
			return sorted[i] < sorted[j]
		})
		tmp := string(sorted)
		buffer[tmp] = append(buffer[tmp], i)
	}
	result := make(map[string][]string)
	for _, words := range buffer {
		if len(words) > 1 {
			sort.Strings(words)
			result[words[0]] = words
		}
	}
	return result
}
