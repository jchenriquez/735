package main

import (
 "fmt"
 "math"
)

func asteroidCollision(asteroids []int) []int {
 stack := make([]int, 0, len(asteroids))

 for _, asteroid := range asteroids {
  stack = append(stack, asteroid)

  for len(stack) > 1 && stack[len(stack)-1] < 0 && stack[len(stack)-2] > 0 {
   righSize := int(math.Abs(float64(stack[len(stack)-1])))
   leftSize := stack[len(stack)-2]

   if righSize > leftSize {
    stack[len(stack)-2], stack[len(stack)-1] = stack[len(stack)-1], stack[len(stack)-2]
    stack = stack[:len(stack)-1]
   } else if leftSize > righSize {
    stack = stack[:len(stack)-1]
   } else {
    stack = stack[:len(stack)-2]
   }
  }
 }

 return stack
}

func main() {
 fmt.Println(asteroidCollision([]int{20, 10, -15}))
}
