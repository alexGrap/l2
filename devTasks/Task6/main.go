package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	fields    string
	delimiter string
	separated bool
}

func main() {
	config := NewConfig()
	cutRun(config)
}

func NewConfig() *Config {
	config := Config{}
	flag.StringVar(&config.fields, "f", "", "List of fields to cut")
	flag.StringVar(&config.delimiter, "d", "\t", "Set custom delimeter")
	flagS := flag.Bool("s", false, "Get only separated strings")
	flag.Parse()
	config.separated = *flagS
	return &config
}

func cut(row string, config *Config) (string, error) {
	var res strings.Builder
	fields := make(map[int]bool)
	var delimeter string
	if config.delimiter != "\t" {
		if len(config.delimiter) == 1 {
			delimeter = config.delimiter
		} else {
			return "", fmt.Errorf("you could set only one character for delimeter")
		}
	}
	if config.fields != "" {
		rangeR := strings.Split(config.fields, ",")
		for _, dPart := range rangeR {
			dRange := strings.Split(strings.TrimSpace(dPart), "-")
			if len(dRange) == 2 {
				dLeft, err := strconv.Atoi(dRange[0])
				if err != nil {
					return "", fmt.Errorf("invalid left value %s", dLeft)
				}
				dRight, err := strconv.Atoi(dRange[1])
				if err != nil {
					return "", fmt.Errorf("invalid right value %s", dRight)
				}
				if dLeft < 1 || dLeft > dRight {
					return "", fmt.Errorf("your range has started from 0 or left border more than right border")
				}
				for i := dLeft; i <= dRight; i++ {
					fields[i] = true
				}
			} else {
				numOfField, err := strconv.Atoi(strings.TrimSpace(dPart))
				if err != nil {
					return "", fmt.Errorf("invalid field value %s", dPart)
				}
				fields[numOfField] = true
			}
		}
	}
	sliceOfRows := strings.Split(row, delimeter)
	if len(sliceOfRows) == 1 && config.separated {
		return "", nil
	}
	isDelim := false
	for i, val := range sliceOfRows {
		_, ok := fields[i+1]
		if ok {
			if isDelim {
				res.WriteString(delimeter + val)
			} else {
				res.WriteString(val)
				isDelim = true
			}
		}
	}
	return res.String(), nil
}

func cutRun(config *Config) {
	var str strings.Builder
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str.WriteString(scanner.Text())
	}
	result, err := cut(str.String(), config)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(result)
}
