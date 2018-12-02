package main

import (
	"bufio"
	"os"
	"strconv"
)

func freqInResults(freq int, results []int) bool {
	for _, res := range results{
		if res == freq {
			return true
		}
	}
	return false
}

func main() {
	println("first advent of code print")
	file, err := os.Open("input.txt")
	if err != nil {
		println(err)
	}
	defer file.Close()

	freq := 0
	scanner := bufio.NewScanner(file)
	results := []int{}
	inputs := []int{}
	for scanner.Scan() {
		delta, err := strconv.Atoi(scanner.Text())
		if err != nil {
			println(err)
		}
		freq += delta
		inputs = append(inputs, delta)
		results = append(results, freq)
	}
	if err := scanner.Err(); err != nil {
		println(err)
	}
	println("result is: ", freq)

	found := false
	for {
		for _, delta := range inputs {
			freq += delta
			if freqInResults(freq, results) {
				println("first repeating frequency: ", freq)
				found = true
				break
			}
		}
		if found {
			break
		}
	}
}


