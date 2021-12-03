package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func getData() []int {
	data, err := ioutil.ReadFile("data")
	if err != nil {
		panic(err)
	}
	datastrings := strings.Split(string(data), "\n")
	var dataset []int

	for _, val := range datastrings {
		x, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		dataset = append(dataset, x)
	}
	return dataset
}

func TestSonarSweep(t *testing.T) {
	var tests = []struct {
		depth    []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}, 7},
		{getData(), 1215},
	}

	for _, test := range tests {
		result := sonarSweep(test.depth)
		require.Equal(t, test.expected, result)
		fmt.Println(result)
	}
}
