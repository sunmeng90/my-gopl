package main

import "fmt"

type student struct {
	name string
}

func (s *student) getName()  {
	fmt.Println("hello world!")
}


func main(){
	var s *student
	s.getName()
	s = nil
	s.getName()
	(*student).getName(nil)
}

/////////////////output////////////////////
//hello world
//hello world
//hello world