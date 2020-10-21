package extract

import (
  "strconv"
)

func getFloatScore(line []string, campo int) float64 {
  i, _ :=  strconv.ParseFloat(line[campo], 64)
  return i
}

func getIntValue(line []string, campo int) int {
  i, _ :=  strconv.Atoi(line[campo])
  return i
}
