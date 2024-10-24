package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

func uniq(input io.Reader, output io.Writer, count bool, onlyDup bool, onlyUniq bool, ignoreCase bool, numFields, numChars int) {
	scanner := bufio.NewScanner(input)
	lines := make(map[string]int)
	order := []string{}

	for scanner.Scan() {
		line := scanner.Text()

		if numFields > 0 {
			words := strings.Fields(line)
			if len(words) > numFields {
				line = strings.Join(words[numFields:], " ")
			} else {
				line = ""
			}
		}

		if numChars > 0 && len(line) > numChars {
			line = line[numChars:]
		}

		if ignoreCase {
			line = strings.Map(unicode.ToLower, line)
		}

		if _, found := lines[line]; !found {
			order = append(order, line)
		}
		lines[line]++
	}

	for _, line := range order {
		countVal := lines[line]

		if (onlyDup && countVal > 1) || (onlyUniq && countVal == 1) || (!onlyDup && !onlyUniq) {
			if count {
				fmt.Fprintf(output, "%d %s\n", countVal, line)
			} else {
				fmt.Fprintln(output, line)
			}
		}
	}
}

func main() {

	count := flag.Bool("c", false, "repetition")
	onlyDup := flag.Bool("d", false, "duplicate str")
	onlyUniq := flag.Bool("u", false, "unique str")
	ignoreCase := flag.Bool("i", false, "ignore")
	numFields := flag.Int("f", 0, "ignore first num_fields")
	numChars := flag.Int("s", 0, "ignore first num_chars")
	flag.Parse()

	if (*count && *onlyDup) || (*count && *onlyUniq) || (*onlyDup && *onlyUniq) {
		fmt.Fprintln(os.Stderr, "incorrect usage -c, -d и -u")
		flag.Usage()
		os.Exit(1)
	}

	var inputFile *os.File
	var outputFile *os.File
	var err error

	switch len(flag.Args()) {
	case 0:
		inputFile = os.Stdin
		outputFile = os.Stdout
	case 1:
		inputFile, err = os.Open(flag.Args()[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error open file %v\n", err)
			os.Exit(1)
		}
		defer inputFile.Close()
		outputFile = os.Stdout
	case 2: // input_file и output_file
		inputFile, err = os.Open(flag.Args()[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error open file %v\n", err)
			os.Exit(1)
		}
		defer inputFile.Close()
		outputFile, err = os.Create(flag.Args()[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error creature file %v\n", err)
			os.Exit(1)
		}
		defer outputFile.Close()
	default:
		fmt.Fprintln(os.Stderr, "a lot of argm")
		flag.Usage()
		os.Exit(1)
	}

	uniq(inputFile, outputFile, *count, *onlyDup, *onlyUniq, *ignoreCase, *numFields, *numChars)
}
