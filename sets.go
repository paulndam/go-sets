package main

import "fmt"

type Set struct {
	integerMap map[int]bool
}

//creates map integer and bool.

func (set *Set) New() {
	set.integerMap = make(map[int]bool)
}

// adds elements if element isn't already in set.
func (set *Set) AddElement(el int) {
	if !set.ContainsElement(el) {
		set.integerMap[el] = true
	}
}

//deletes elements
func (set *Set) DeleteElement(el int) {
	delete(set.integerMap, el)
}

func (set *Set) ContainsElement(el int) bool {
	var exist bool
	_, exist = set.integerMap[el]
	return exist
}


func main(){
	var s *Set 
	s = &Set{}
	s.New()
	s.AddElement(1)
	s.AddElement(2)
	s.AddElement(3)
	s.AddElement(4)
	s.AddElement(5)

	fmt.Println(s)
	fmt.Println(s.ContainsElement(1))
}
