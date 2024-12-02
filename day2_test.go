package aoc_test

import (
	"testing"

	"github.com/jeynesyra/aoc-2024"
)

func TestDay2(t *testing.T) {
	rs := []*aoc.Report{
		aoc.NewReport([]int{7, 6, 4, 2, 1}),
		aoc.NewReport([]int{1, 2, 7, 8, 9}),
		aoc.NewReport([]int{9, 7, 6, 2, 1}),
		aoc.NewReport([]int{1, 3, 2, 4, 5}),
		aoc.NewReport([]int{8, 6, 4, 4, 1}),
		aoc.NewReport([]int{1, 3, 6, 7, 9}),
	}

	task1Count := 0
	for _, r := range rs {
		if r.CheckSafety() {
			task1Count++
		}
	}

	task2Count := 0
	for _, r := range rs {
		if r.CheckSafetyWithProblemDampener() {
			task2Count++
		}
	}

	if task1Count != 2 {
		t.Errorf("expected %d safe reports, got %d", 2, task1Count)
	}
	if task2Count != 4 {
		t.Errorf("expected %d safe reports, got %d", 4, task2Count)
	}
}
