package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func makeIntcodes(inputArr []string) []int {
  intArr := make([]int, len(inputArr))
  for i, v := range inputArr {
    intArr[i], _ = strconv.Atoi(v)
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
		inputStrArr := strings.Split(lineStr, ",")
    
    for noun := range make([]int, 99) {
      for verb := range make([]int, 99) {
        ic := makeIntcodes(inputStrArr)
        ic[1] = noun
        ic[2] = verb
        pos := 0
        for {
          opcode := ic[pos]
          i1, i2, i3 := ic[pos+1], ic[pos+2], ic[pos+3]
          if opcode == 1 {
            ic[i3] = ic[i1] + ic[i2]
          } else if opcode == 2 {
            ic[i3] = ic[i1] * ic[i2]
          } else {
            break
          }
          pos += 4
        }
        if ic[0] == 19690720 {
          fmt.Println(noun, verb)
        }
      }
    }
	}
}
