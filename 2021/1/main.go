package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	fName := flag.String("filename", "sample.txt", "the file to read to get the input from")
	flag.Parse()

	input, err := readFile(*fName)
	if err != nil {
		fmt.Printf("err: %+v\n", err)
		os.Exit(1)
	}

	fmt.Println(increasesPart1(input))
	fmt.Println(increasesPart2(input))
}

// increasesPart1 AOC part 1
func increasesPart1(input []int) int {
	increases := 0
	last := 0

	for i, x := range input {
		if i == 0 {
			last = x
			continue
		}

		if x > last {
			increases++
		}
		last = x
	}

	return increases
}

func increasesPart2(input []int) int {
	var preprocessed []int

	for i := 0; i < len(input)-2; i++ {
		sum := input[i] + input[i+1] + input[i+2]
		preprocessed = append(preprocessed, sum)
	}

	return increasesPart1(preprocessed)
}

func readFile(fname string) ([]int, error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")
	// Assign cap to avoid resize on every append.
	nums := make([]int, 0, len(lines))

	for _, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 {
			continue
		}
		// Atoi better suits the job when we know exactly what we're dealing
		// with. Scanf is the more general option.
		n, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}

	return nums, nil
}