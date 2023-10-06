package main 

func numPairs(values []int) int {
  var result int 
  var counter map[int]int = {}

  for _, value := range values {
    currCount, ok = counter[value]
    if !ok {
      counter[value] = 0    
    }
    result += counter[value]
    counter[value] += 1
  }

  return result
  
}
