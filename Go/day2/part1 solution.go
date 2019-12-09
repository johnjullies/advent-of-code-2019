package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
  "strconv"
)

func runOpcodes(ic []int) []int {
  for i := 0; true; i += 4 {
    opcode := ic[i]
    i1, i2, i3 := ic[i+1], ic[i+2], ic[i+3]
    if opcode == 1 {
      ic[i3] = ic[i1] + ic[i2]
    } else if opcode == 2 {
      ic[i3] = ic[i1] * ic[i2]
    } else {
      break
    }
  }
  return ic
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
    intArr[1] = 12
    intArr[2] = 2
		fmt.Println(runOpcodes(intArr))
	}
}
