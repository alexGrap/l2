package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
)

type Config struct {
	after       int
	before      int
	contextRows int
	count       bool
	ignoreCase  bool
	invert      bool
	fixed       bool
	strNum      bool
	reg         string
	filename    string
}

func main() {
	config := NewConfig()
	res, err := startGrep(config)
	if err != nil {
		log.Fatalf(err.Error())
	}
	switch result := res.(type) {
	case []string:
		for _, row := range result {
			fmt.Println(row)
		}
	case []int:
		for _, row := range result {
			fmt.Println(row)
		}
	case string:
		fmt.Println(result)
	case int:
		fmt.Println(result)
	default:
		fmt.Printf("Unknowed type %T\n", result)
	}
}

func NewConfig() *Config {
	config := Config{}
	flag.IntVar(&config.after, "A", 0, "show +N strings after match")
	flag.IntVar(&config.before, "B", 0, "show +N strings before match")
	flag.IntVar(&config.contextRows, "C", 0, "(A+B) show ±N strings around match")
	flagC := flag.Bool("c", false, "numbers of strings")
	flagI := flag.Bool("i", false, "ignore case")
	flagV := flag.Bool("v", false, "instead of match, exclude")
	flagF := flag.Bool("f", false, "exact string match, not a pattern")
	flagN := flag.Bool("n", false, "print line number")

	flag.Parse()

	args := flag.Args()
	config.count = *flagC
	config.ignoreCase = *flagI
	config.invert = *flagV
	config.fixed = *flagF
	config.strNum = *flagN

	if len(args) == 2 {
		config.reg = args[0]
		config.filename = args[1]
	} else {
		log.Fatalf("more than 1 file can not be going")
	}
	return &config
}

func fileReader(filename string) ([]string, error) {
	var rows []string
	file, err := os.Open(filename)
	if err != nil {
		return rows, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}
	return rows, nil
}

func grep(rows []string, config *Config) (interface{}, error) {
	var prefix, postfix string
	if config.ignoreCase {
		prefix = "(?i)"
	}
	if config.fixed {
		prefix = "^"
		postfix = "$"
	}
	regex, err := regexp.Compile(prefix + config.reg + postfix)
	if err != nil {
		return "Error", fmt.Errorf("uncorrected regular expression")
	}

	switch {
	case config.after != 0:
		for i, row := range rows {
			if regex.MatchString(row) {
				if config.after <= len(rows)-i {
					return rows[i : i+config.after+1], nil
				}
				return rows[i:], nil
			}
		}
		return "not found", nil
	case config.before != 0:
		for i, row := range rows {
			if regex.MatchString(row) {
				if config.before-1 <= i {
					return rows[i-config.before : i+1], nil
				}
				return rows[:i+1], nil
			}
		}
		return "not found", nil
	case config.contextRows != 0:
		for i, row := range rows {
			if regex.MatchString(row) {
				firstI, lastI := 0, len(rows)
				if config.contextRows-1 <= i {
					firstI = i - config.contextRows
				}
				if config.contextRows <= len(rows)-i {
					lastI = i + config.contextRows + 1
				}
				return rows[firstI:lastI], nil
			}
		}
		return "not found", nil
	case config.count:
		total := 0
		for _, row := range rows {
			total += len(regex.FindAllString(row, -1))
		}
		if config.invert {
			return len(rows) - total, nil
		}
		return total, nil
	case config.strNum:
		var numOfRows []int
		for i, row := range rows {
			if regex.MatchString(row) {
				numOfRows = append(numOfRows, i)
			}
		}
		return numOfRows, nil
	default:
		var res []string
		for _, row := range rows {
			if regex.MatchString(row) {
				res = append(res, row)
			}
		}
		return res, nil
	}
}

func startGrep(config *Config) (interface{}, error) {
	rows, err := fileReader(config.filename)
	if err != nil {
		return "", fmt.Errorf("can not read file%s: %s", config.filename, err.Error())
	}
	return grep(rows, config)
}
