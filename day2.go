package aoc

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/jeynesyra/aoc-2024/pkg"
)

func Day2() error {
	lines, err := pkg.ReadFile("data/2.txt")
	if err != nil {
		return err
	}

	var safetyScore int
	var reports []*Report
	for _, line := range lines {
		levels := strings.Split(line, " ")

		var lvls []int
		for _, level := range levels {
			lvl, err := strconv.Atoi(level)
			if err != nil {
				fmt.Printf("error converting '%s': %v\n", level, err)
				continue
			}
			lvls = append(lvls, lvl)
		}

		r := NewReport(lvls)
		if r.CheckSafety() {
			safetyScore++
		}

		reports = append(reports, r)
	}

	fmt.Println("_____ Day 1 _____")
	fmt.Printf("Task 1: %d\n", safetyScore)

	safetyScore = 0
	for _, r := range reports {
		if r.CheckSafetyWithProblemDampener() {
			safetyScore++
		}
	}

	fmt.Printf("Task 2: %d\n", safetyScore)

	return nil
}

type Report struct {
	Levels []int
}

func NewReport(levels []int) *Report {
	return &Report{
		Levels: levels,
	}
}

func (r *Report) CheckSafety() bool {
	incChk := true
	decChk := true
	diffChk := true

	// Check all increasing
	for i := 1; i < len(r.Levels); i++ {
		if r.Levels[i] <= r.Levels[i-1] {
			incChk = false
			break
		}
	}

	// Check decreasing
	for i := 1; i < len(r.Levels); i++ {
		if r.Levels[i] >= r.Levels[i-1] {
			decChk = false
			break
		}
	}

	// Check safety
	for i := 1; i < len(r.Levels); i++ {
		diff := math.Abs(float64(r.Levels[i] - r.Levels[i-1]))
		if diff < 1 || diff > 3 {
			diffChk = false
		}
	}

	return (incChk || decChk) && diffChk
}

func (r *Report) CheckSafetyWithProblemDampener() bool {
	if r.CheckSafety() {
		return true
	}

	original := make([]int, len(r.Levels))
	copy(original, r.Levels)

	for i := 0; i < len(r.Levels); i++ {
		newLevels := append([]int{}, r.Levels[:i]...)    // Create a new slice from the part before i
		newLevels = append(newLevels, r.Levels[i+1:]...) // Append the part after i
		r.Levels = newLevels
		if r.CheckSafety() {
			return true
		}

		r.Levels = original
	}

	return false
}
