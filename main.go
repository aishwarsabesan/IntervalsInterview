package main

import (
  "fmt"
)

type interval struct {
  startDate string
  endDate string
}

func main() {
  
  existingSet := []interval{
    {
      startDate: "2020-01-01",
      endDate: "2020-01-05",
    },
    {
      startDate: "2020-01-10",
      endDate: "2020-01-20",
    },
  }
  
  incomingInterval := interval{
    startDate: "2020-01-06",
    endDate: "2020-01-09",
  }
  
  overlaps := checkOverlap(incomingInterval, existingSet)

  fmt.Println(overlaps)
}

func checkOverlap(incoming interval, existing []interval) bool {
  // IMPLEMENT ME
  // FUNCTION SHOULD RETURN TRUE IF INCOMING INTERVAL OVERLAPS ANY OF THE EXISTING INTERVALS
  // IT SHOULD ALSO RETURN TRUE IF INCOMING INTERVAL FULLY ENCAPSULATES ONE ORE MORE OF THE EXISTING INTERVALS
  
  return false
}
