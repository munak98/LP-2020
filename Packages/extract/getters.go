package extract

import (
  "strconv"
)

func fromUF(line []string, UF string) bool{
    if line[5] == UF{
      return true
    } else{
      return false
    }
}

func getScoreCN(line []string) float64 {
    i, _ :=  strconv.ParseFloat(line[91], 64)
    return i
}

func getScoreCH(line []string) float64 {
    i, _ :=  strconv.ParseFloat(line[92], 64)
    return i
}

func getScoreLC(line []string) float64 {
    i, _ :=  strconv.ParseFloat(line[93], 64)
    return i
}

func getScoreMT(line []string) float64 {
    i, _ :=  strconv.ParseFloat(line[94], 64)
    return i
}
