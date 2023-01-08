package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day1/input.txt")

	if err != nil {
		fmt.Println(err)
	}

	var mostCalories int64
	var currentReindeer = 1
	var calories = int64(0)

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		if fileScanner.Text() == "" {
			if calories > mostCalories {
				mostCalories = calories
			}

			currentReindeer++
			calories = 0
			continue
		}

		var number, _ = strconv.ParseInt(fileScanner.Text(), 0, 0)
		calories += number
	}

	fmt.Printf("%v", mostCalories)
}
