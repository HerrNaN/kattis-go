package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Scenario struct {
	heads, knights []int
}

type Input []Scenario

func main() {
	reader := bufio.NewReader(os.Stdin)
	scenarios := readInput(reader)
	run(scenarios)
}

func run(scenarios Input) {
	for _, s := range scenarios {
		sort.Ints(s.heads)
		sort.Ints(s.knights)
		cost, err := slayCost(s.heads, s.knights)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(cost)
		}
	}
}

func readInput(r *bufio.Reader) Input {
	var scenarios []Scenario
	for {
		nh, nk := readHeader(r)
		if nh == 0 && nk == 0 {
			break
		}

		hs, ks := readScenario(r, nh, nk)
		scenarios = append(scenarios, Scenario{hs, ks})
	}

	return scenarios
}

func readHeader(r *bufio.Reader) (int, int) {
	var heads, knights int
	fmt.Fscanf(r, "%d %d\n", &heads, &knights)

	return heads, knights
}

func readScenario(r *bufio.Reader, nh, nk int) ([]int, []int) {
	var hs []int
	for i := 0; i < nh; i++ {
		var h int
		fmt.Fscanf(r, "%d\n", &h)
		hs = append(hs, h)
	}
	var ks []int
	for i := 0; i < nk; i++ {
		var k int
		fmt.Fscanf(r, "%d\n", &k)
		ks = append(ks, k)
	}
	return hs, ks
}

// Needs to be sorted
func slayCost(heads []int, knights []int) (int, error) {
	if len(heads) == 0 && len(knights) == 0 {
		return 0, nil
	}
	if len(knights) == 0 {
		return 0, fmt.Errorf("Loowater is doomed!")
	}
	if len(heads) == 0 {
		return 0, nil
	}
	if len(heads) > len(knights) {
		return 0, fmt.Errorf("Loowater is doomed!")
	}
	if heads[0] > knights[0] {
		return slayCost(heads, knights[1:])
	}

	cost, err := slayCost(heads[1:], knights[1:])
	if err != nil {
		return 0, err
	}
	return knights[0] + cost, nil
}
