package main

import (
  "fmt"
  "strings"
  "bufio"
  "os"
  "strconv"
  "math"
)

func haveSameSign(a, b int) bool {
  return a * b > 0
}

func main() {
  file, err := os.Open("input.txt")
  if err != nil {
    fmt.Println(err)
  }
  defer file.Close()
  
  total_safe := 0

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    values := strings.Fields(line)
    
    first, err1 := strconv.Atoi(values[0])
    second, err2 := strconv.Atoi(values[1])
    if err1 != nil || err2 != nil {
      fmt.Println("Error converting to int")
    }

    is_level_safe := true

    prev_diff := first - second
    abs_prev_diff := math.Abs(float64(prev_diff))

    if !(abs_prev_diff >= 1 && abs_prev_diff <= 3) {
      continue
    }

    for i := 2; i < len(values); i++ {
      prev, err1 := strconv.Atoi(values[i-1])
      current, err2 := strconv.Atoi(values[i])

      if err1 != nil || err2 != nil {
        fmt.Println("Error converting to int")
      }

      diff := prev - current
      abs_diff := math.Abs(float64(diff))
      if !(haveSameSign(prev_diff, diff) && abs_diff >= 1 && abs_diff <= 3) {
        is_level_safe = false
      }
    }

    if is_level_safe {
      total_safe++
    }
  }

  fmt.Println(total_safe)
}
