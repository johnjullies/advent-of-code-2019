package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
  "strconv"
)

func runOpcodes(intArr []int) []int {
  for i := 0; i < len(intArr); i += 4 {
    switch intArr[i] {
      case 1:
        val1, val2, val3 := intArr[i+1], intArr[i+2], intArr[i+3]
        intArr[val3] = intArr[val1] + intArr[val2]
      case 2:
        val1, val2, val3 := intArr[i+1], intArr[i+2], intArr[i+3]
        intArr[val3] = intArr[val1] * intArr[val2]
      default:
        return intArr
    }
  }
  return intArr
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lineStr := scanner.Text()
		inputArr := strings.Split(lineStr, ",")
    intArr := make([]int, len(inputArr))
    for i, v := range inputArr {
      intArr[i], _ = strconv.Atoi(v)
    }

		fmt.Println(runOpcodes(intArr))
	}
}
