package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Config struct {
	column   int
	byDigit  bool
	reversed bool
	unique   bool
	filename string
}

func main() {
	s := createConfig()
	res, err := actionStart(s)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(res)
}

func createConfig() *Config {
	s := Config{}
	flag.IntVar(&s.column, "k", 0, "column for sort")
	flagN := flag.Bool("n", false, "sort by digit")
	flagR := flag.Bool("r", false, "sort in reverse direction")
	flagU := flag.Bool("u", false, "show only unique values")
	flag.Parse()
	args := flag.Args()
	s.byDigit = *flagN
	s.reversed = *flagR
	s.unique = *flagU

	if len(args) == 1 {
		s.filename = args[0]
	} else {
		log.Fatalf("not enough argument")
	}
	return &s
}

func readFile(filename string) ([]string, error) {
	var rows []string
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return rows, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}
	return rows, nil
}

func getUnique(rows []string) []string {
	uniqueMap := make(map[string]bool)
	for _, row := range rows {
		uniqueMap[row] = true
	}
	uniqueRows := rows[:0]
	for key := range uniqueMap {
		uniqueRows = append(uniqueRows, key)
	}
	return uniqueRows
}

func actionStart(s *Config) (string, error) {
	rows, err := readFile(s.filename)
	if err != nil {
		return "", fmt.Errorf("file doesn't exist %s\n, %s", s.filename, err.Error())
	}
	return sortRows(rows, s)
}

func getColumn(row string, s *Config) (string, error) {
	re := regexp.MustCompile(`\s+`)
	columnsSlice := re.Split(strings.TrimSpace(row), -1)
	if len(columnsSlice) >= s.column {
		return columnsSlice[s.column-1], nil
	}
	return "", fmt.Errorf("can't find column")
}

func sortRows(rows []string, s *Config) (string, error) {
	switch {
	case s.column > 0:
		if s.unique {
			rows = getUnique(rows)
		}
		sort.SliceStable(rows, func(i, j int) bool {
			ix, err := getColumn(rows[i], s)
			if err != nil {
				return false
			}
			jx, err := getColumn(rows[j], s)
			if err != nil {
				return false
			}
			if s.byDigit {
				if s.reversed {
					return !comparator(ix, jx)
				}
				return comparator(ix, jx)
			}
			if s.reversed {
				return ix < jx
			}
			return ix > jx
		})
	case s.unique:
		rows = getUnique(rows)
		sort.SliceStable(rows, func(i, j int) bool {
			if s.reversed {
				return rows[i] > rows[j]
			}
			return rows[i] < rows[j]
		})
	case s.byDigit:
		if s.unique {
			rows = getUnique(rows)
		}
		sort.SliceStable(rows, func(i, j int) bool {
			if s.reversed {
				return !comparator(rows[i], rows[j])
			}
			return comparator(rows[i], rows[j])
		})

	default:
		sort.SliceStable(rows, func(i, j int) bool {
			if s.reversed {
				return rows[i] > rows[j]
			}
			return rows[i] < rows[j]
		})
	}

	var res strings.Builder
	rowLength := len(rows)
	for i, row := range rows {
		if i < rowLength-1 {
			_, _ = res.WriteString(row + "\n")
		} else {
			_, _ = res.WriteString(row)
		}
	}
	return res.String(), nil
}

func comparator(strI, strJ string) bool {
	i, _ := strconv.Atoi(strI)
	j, _ := strconv.Atoi(strJ)
	return i < j
}
