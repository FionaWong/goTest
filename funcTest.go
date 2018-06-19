package main

import (
  "fmt"
  "math/rand"
)


func main() {
  // testFor()
  // testIf()
  // testRange()
  fmt.Println(testCase())
}

func testIf (){
  x := 100;
  if x > 10 {
    fmt.Println("x is valued more than 10")
  }else if x>100 {
      fmt.Println("x is valued more than 10")
  }else{
    fmt.Println("100")
  }
}

func testFor () {
  sum := 0
  for i:=10;i<100;i++ {
    if sum >50 {
      break
    }
    sum += i
  }
  fmt.Println("sum is ",sum)
}

func testRange () {
  numbers := make(map[string] int )
  numbers["one"] = 1
  numbers["two"] = 2
  numbers["three"] = 3

  for k,v := range numbers {
    fmt.Printf("%s value is %d", k, v )
  }
}

func testCase () int{
  s1 := rand.NewSource(42)
  r1 := rand.New(s1)
  x := r1.Intn(30)
  switch x {
    case 9 :
      fmt.Println("x is more than 90")
      fallthrough
    case 8 :
      fmt.Println("x is more than 80")
      fallthrough
    case 7 :
      fmt.Println("x is more than 70")
      fallthrough
    default:
      fmt.Println("x is less than 70")

  }
  return x
}
