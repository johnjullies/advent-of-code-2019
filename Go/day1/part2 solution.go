package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func addFuel(fuel int) int {
  finalFuel := 0
  
  for fuel > 0 {
    finalFuel += fuel
    fuel = fuel / 3.0
    roundDown := math.Round(float64(fuel))
    fuel = int(roundDown) - 2
  } 

  return finalFuel
}

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
    finalFuel := addFuel(fuelReq)
    total += finalFuel
	}

	fmt.Println(total)
}
