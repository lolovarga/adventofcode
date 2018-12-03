package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func parse_claim(claim string) []int {
	//#1 @ 12,548: 19x10
	claimID := strings.Split(claim, "@")[0][1:]
	parsedClaimID, err := strconv.Atoi(strings.TrimSpace(claimID))
	if err != nil {
		println(err)
	}
	//println("claimId:", parsedClaimID)
	coordinates := strings.Split(strings.Split(claim, "@")[1], ":")
	relativeCoordinates := coordinates[0]
	matrixSizes := coordinates[1]
	//println("coords:", relativeCoordinates)
	//println("matrix", matrixSizes)
	fromLeft, err := strconv.Atoi(strings.TrimSpace(strings.Split(relativeCoordinates, ",")[0]))
	if err != nil {
		println(err)
	}
	//println("fromLeft", fromLeft)
	fromTop, err := strconv.Atoi(strings.TrimSpace(strings.Split(relativeCoordinates, ",")[1]))
	if err != nil {
		println(err)
	}
	//println("fromLeft", fromTop)
	dx, err := strconv.Atoi(strings.TrimSpace(strings.Split(matrixSizes, "x")[0]))
	if err != nil {
		println(err)
	}
	//println("dx", dx)
	dy, err := strconv.Atoi(strings.TrimSpace(strings.Split(matrixSizes, "x")[1]))
	if err != nil {
		println(err)
	}
	//println("dy", dy)
	parsedClaim := []int{parsedClaimID, fromLeft, fromTop, dx, dy}
	return parsedClaim
}

func main() {
	println("test")
	file, err := os.Open("./input.txt")
	if err != nil {
		println(err)
	}
	println(file)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var inputs [][]int
	fabric := [1000][1000]string{}
	for scanner.Scan() {
		claim := scanner.Text()
		parsedClaim := parse_claim(claim)
		inputs = append(inputs, parsedClaim)
	}
	overClaimed := 0
	for _, claim := range inputs {
		for x := 0; x < claim[3]; x++ {
			for y := 0; y < claim[4]; y++ {
				if fabric[x + claim[1]][y + claim[2]] == "" {
					fabric[x + claim[1]][y + claim[2]] = strconv.Itoa(claim[0])
				} else {
					if fabric[x + claim[1]][y + claim[2]] != "X" {
						overClaimed++
						fabric[x + claim[1]][y + claim[2]] = "X"
					}
				}

			}
		}
	}
	println("overClaimed: ", overClaimed)

	for _, claim := range inputs {
		inches := 0
		for x := 0; x < claim[3]; x++ {
			for y := 0; y < claim[4]; y++ {
				if fabric[x + claim[1]][y + claim[2]] == strconv.Itoa(claim[0]) {
					inches++
				}
			}
		}
		if inches == claim[3] * claim[4] {
			println("this is it: ", claim[0])
		}
	}
}