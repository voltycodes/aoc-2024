package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"
  "strconv"
  "sort"
)

func main()  {
  file, err := os.Open("input.txt")
  if err != nil {
    fmt.Println("Error reading file")
    return
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  // make two arrays for the two values
  var slice1 []int
  var slice2 []int

  for scanner.Scan() {
    line := scanner.Text()
    values := strings.Fields(line)

    if len(values) != 2 {
      fmt.Println("Expected 2 values, got", len(values))
      continue
    }

    first, err1 := strconv.Atoi(values[0])
    second, err2 := strconv.Atoi(values[1])

    if err1 != nil || err2 != nil {
      fmt.Println("Error converting string values to integers")
      continue
    }
    
    slice1 = append(slice1, first)
    slice2 = append(slice2, second)
  }

  // sort the slices
  sort.Ints(slice1)
  sort.Ints(slice2)

  total := 0
  length := len(slice1)
  
  for i := 0; i < length; i++ {
    diff := slice2[i] - slice1[i]
    if diff > 0 {
      total += diff
    } else {
      total += -diff
    }
  }

  fmt.Println(total)
}
