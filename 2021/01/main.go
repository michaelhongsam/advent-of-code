package main

func sonarSweep(depth []int) int {
	count := 0
	var prev int
	for i, val := range depth {
		if i != 0 && val > prev {
			count += 1
		}
		prev = val
	}
	return count
}
