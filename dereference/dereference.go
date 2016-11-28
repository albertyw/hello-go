package main

import "fmt"

type Obj struct {
	ID      int
	Name    string
	RefIDs  []int
	RefObjs []Obj
}

var obj1 = Obj{1, "Joe", []int{20, 21}, make([]Obj, 2)}
var obj2 = Obj{20, "Harvard", []int{80}, make([]Obj, 1)}
var obj3 = Obj{21, "MIT", []int{80}, make([]Obj, 1)}
var obj4 = Obj{80, "Boston", []int{}, make([]Obj, 0)}

func main() {
	data := []*Obj{&obj1, &obj2, &obj3, &obj4}
	fmt.Println(dereference(&obj1, data))
}

func dereference(seed *Obj, data []*Obj) *Obj {
	// TODO
	return seed
}
