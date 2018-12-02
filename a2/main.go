package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	inputs := []string{}
	checksum := 0
	twice := 0
	treeTimes := 0

	for scanner.Scan() {
		boxID := scanner.Text()
		inputs = append(inputs, scanner.Text())
		found2 := false
		found3 := false
		for _, char := range boxID {
			if strings.Count(boxID, string(char)) == 3 {
				if !found3 {
					treeTimes++
					found3 = true
				}
			}
			if strings.Count(boxID, string(char)) == 2 {
				if !found2 {
					twice++
					found2 = true
				}
			}
		}
	}
	checksum = twice * treeTimes
	println("checksum is: ", checksum)

	box1 := ""
	box2 := ""

	for index1, boxID1 := range inputs {
		for index2, boxID2 := range inputs {
			differences := 0
			if index1 != index2 {
				for charIndex, _ := range boxID1 {
					if boxID1[charIndex] != boxID2[charIndex]{
						differences++
					}
					if differences > 1 {
						break
					}
				}
				if differences > 1 {
					continue
				}
				println(boxID1, index1, boxID2, index2)
				box1 = boxID1
				box2 = boxID2
			}
		}
	}
	println(box1, box2)
	result := ""
	for i, _ := range box1 {
		if box1[i] == box2[i] {
			result += string(box1[i])
		}
	}
	println("result: ", result)
}