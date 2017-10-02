package main

import (
	"fmt"
	"sort"
	//"strings"
	"math/rand"
)

type Person struct {
	ID        int
	FirstName string
	LastName  string
}

type PersonSlice []*Person

func (p PersonSlice) Len() int { return len(p) }
func (p PersonSlice) Swap(i,j int) { p[i], p[j] = p[j], p[i]}
func (p PersonSlice) Less(i,j int) bool {
	// LastName, then FirstName, then ID
	if p[i].LastName < p[j].LastName {
		return true
	} else if p[i].LastName == p[j].LastName {
		if p[i].FirstName < p[j].FirstName {
			return true
		} else if p[i].FirstName == p[j].FirstName {
			return p[i].ID < p[j].ID
		} else { return false }
	} else { return false }
}
func (p Person) String() string {
	return fmt.Sprintf("%s, %s : %d", p.LastName, p.FirstName, p.ID)
}

func NewPerson(first, last string) *Person {
	// TODO
	p := new(Person)
	p.FirstName = first
	p.LastName = last
	p.ID = rand.Intn(1000000)
	return p
}

func main() {
	var people PersonSlice
	p1 := NewPerson("Juan", "Perez")
	p2 := NewPerson("Juan", "Soto")
	p3 := NewPerson("Carlos", "Perez")
		
	/*{FirstName:"Juan", LastName:"Soto", ID:rand.Intn(1000000)},
		{FirstName:"Carlos", LastName:"Perez", ID:rand.Intn(1000000)},
	}*/
	people = append(people, p1)
	people = append(people, p2)
	people = append(people, p3)
	//people_slice := people[:]
	fmt.Println(people)
	
	sort.Sort(people)
	fmt.Println(people)



	/*
	for k, v := range hello {
		fmt.Printf("%s, %s\n", string, v)
	}
	*/
}


