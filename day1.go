package aoc

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/jeynesyra/aoc-2024/pkg"
)

// Locations is used to store all locations found by the historians.
type Locations struct {
	a []int
	b []int
}

// NewLocations sorts 2 lists of locations and creates a Locations domain.
func NewLocations(a, b []int) Locations {
	// Sort the slices before creating the domain
	sort.Ints(a)
	sort.Ints(b)

	return Locations{
		a: a, b: b,
	}
}

func (l *Locations) differences() (diff int) {
	for i, id := range l.a {
		d := math.Abs(float64(id - l.b[i]))
		diff += int(d)
	}
	return diff
}

func (l *Locations) similarity() (score int) {
	instances := map[int]int{}

	for _, aid := range l.a {
		for _, bid := range l.b {
			if aid == bid {
				instances[aid]++
			}
		}
	}

	for _, aid := range l.a {
		count := instances[aid]

		score += aid * count
	}

	return score
}

func Day1() error {
	lines, err := pkg.ReadFile("data/1.txt")
	if err != nil {
		return err
	}

	var a, b []int
	for _, line := range lines {
		digits := strings.Fields(line)

		if len(digits) != 2 {
			return fmt.Errorf("expected length of digits to be 2, got %d", len(digits))
		}

		l, err := strconv.Atoi(digits[0])
		if err != nil {
			return err
		}
		a = append(a, l)

		r, err := strconv.Atoi(digits[1])
		if err != nil {
			return err
		}
		b = append(b, r)
	}

	locs := NewLocations(a, b)
	fmt.Println("_____ Day 1 _____")
	fmt.Printf("Task 1: %d\n", locs.differences())
	fmt.Printf("Task 2: %d\n", locs.similarity())

	return nil
}
