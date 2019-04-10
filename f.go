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

func ExpandRange(rng []string) ([]int, error) {
	var expanded_range []int

	for _, r := range rng {
		if strings.Contains(r, "-") {
			ends := strings.Split(r, "-")
			start, err := strconv.Atoi(ends[0])
			if err != nil {
				return nil, fmt.Errorf("invalid range: %s", r)
			}
			end, err := strconv.Atoi(ends[1])
			if err != nil {
				return nil, fmt.Errorf("invalid range: %s", r)
			}
			if start >= end {
				return nil, fmt.Errorf("invalid range: %s", r)
			}
			for i := start; i <= end; i++ {
				expanded_range = append(expanded_range, i)
			}
		} else {
			i, err := strconv.Atoi(r)
			if err != nil {
				return nil, fmt.Errorf("invalid range: %s", r)
			}
			expanded_range = append(expanded_range, i)
		}
	}

	return expanded_range, nil
}

func main() {
	// take care of command line flags
	file := flag.String("file", "", "file to parse")
	sep := flag.String("s", " ", "field output separator")

	flag.Parse()

	// args left over after parsing flags
	// this should be a list of field and field ranges
	args := flag.Args()

	indices, err := ExpandRange(args)

	if err != nil {
		fmt.Fprintf(os.Stderr, "f: %s\n", err)
		os.Exit(1)
	}

	// fd will point at STDIN or file specified
	// by the -file flag
	var fd *os.File
	if *file != "" {
		fd, err = os.Open(*file)
		if err != nil {
			log.Fatal("Could not open file")
		}
		defer fd.Close()
	} else {
		fd = os.Stdin
	}

	// Scan file for fields
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		var output []string

//		output = GetFields(scanner.Text(), args)

		fields := strings.Fields(scanner.Text())

		for _, index := range indices {
			if len(fields) > index-1 {
				output = append(output, fields[index-1])
			}
		}

		// Output fields, joined by string separator
		// specified in -s flag
		fmt.Println(strings.Join(output, *sep))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
