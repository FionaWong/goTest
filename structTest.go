package main

import(
  "fmt"

)

func main() {
  //one
  var P person
  P.name = "tom"
  P.age = 10
  //two
  P2 := person{"jim", 12}
  //three
  P3 := person{name:"john", age:23}

  fmt.Println(older(P2, P3))
}

type person struct {
  name string
  age int
}

func older (p1, p2 person) ( desc string, diffVal int) {
   old := theOlderMan(&p1, &p2)[0]
   young := theOlderMan(&p1, &p2)[1]
  diffVal = (old.age-young.age)
  desc = fmt.Sprintf("%s is older than %s for %d years", old.name, young.name, diffVal)
  return
}

func theOlderMan (p1, p2 *person) ( ret [2] person ) {
  old := *p1
  young := *p2
  if p1.age < p2.age {
    old = *p2
    young = *p1
  }
  ret[0] = old
  ret[1] = young
  return ret
}
