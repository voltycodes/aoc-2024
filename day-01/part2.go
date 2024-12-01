package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

func countOccurences(slice1, slice2 []int) int {
  countMap := make(map[int]int)
  for _, num := range slice2 {
    countMap[num]++
  }

  total := 0
  for _, num := range slice1 {
    total += num * countMap[num]
  }

  return total
}

func main()  {
  file, err := os.Open("input.txt")
  if err != nil {
    fmt.Println(err)
  }
  defer file.Close()

  var slice1 []int
  var slice2 []int

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    values := strings.Fields(line)

    if len(values) != 2 {
      fmt.Println("Invalid input")
      continue
    }

    first, err1 := strconv.Atoi(values[0])
    second, err2 := strconv.Atoi(values[1])

    if err1 != nil || err2 != nil {
      fmt.Println("Invalid input")
      continue
    }

    slice1 = append(slice1, first)
    slice2 = append(slice2, second)
  }

  total := countOccurences(slice1, slice2)
  fmt.Println(total)
}
