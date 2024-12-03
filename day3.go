package aoc

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/jeynesyra/aoc-2024/pkg"
)

func Day3() error {
	lines, err := pkg.ReadFile("data/3.txt")
	if err != nil {
		return err
	}

	var allLines string
	for _, line := range lines {
		allLines += line
	}

	fmt.Println("_____ Day 3 _____")

	t1Res, err := task1(allLines)
	if err != nil {
		return err
	}
	t2Res, err := task2(allLines)
	if err != nil {
		return err
	}

	fmt.Printf("Task 1: %d\n", t1Res)
	fmt.Printf("Task 2: %d\n", t2Res)

	return nil
}

func task1(allLines string) (int64, error) {
	pattern := `mul\(([0-9]{1,3}),([0-9]{1,3})\)`
	re := regexp.MustCompile(pattern)

	matches := re.FindAllStringSubmatch(allLines, -1)

	var result int64
	if len(matches) > 0 {
		for _, match := range matches {
			a, err := strconv.Atoi(match[1])
			if err != nil {
				return 0, err
			}

			b, err := strconv.Atoi(match[2])
			if err != nil {
				return 0, err
			}

			result = result + int64(a*b)
		}
	}

	return result, nil
}

func task2(allLines string) (result int64, err error) {
	// Start at beginning, find all before first don't()
	// Cut then find all instances between do() and don't()
	// Cut then find any with final do()

	return 0, nil
}
