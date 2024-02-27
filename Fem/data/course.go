package data

import "fmt"

type Duration = float64

type Course struct {
	Id       int
	Name     string
	Slug     string
	Legacy   bool
	Duration Duration
}

// implicit implementation of KindOfSignable interface
func (c Course) SignUp() bool {
	return true
}

func (c Course) String() string {
	return fmt.Sprintf("%v - %v - %v", c.Name, c.Slug, c.Duration)
}
