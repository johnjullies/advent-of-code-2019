package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		lineStr := scanner.Text()
		mass, _ := strconv.ParseFloat(lineStr, 64)
		fuel := mass / 3.0
		roundDown := math.Floor(fuel)
		fuelReq := int(roundDown) - 2
		total += fuelReq
	}

	fmt.Println(total)
}
