package data

import "time"

type Workshop struct {
	// embed the Course type now all properties and methods of Course are available to Workshop
	Course
	// can embed more than one type
	Instructor
	Date time.Time
}

// implicit implementation of KindOfSignable interface
func (w Workshop) SignUp() bool {
	return true
}

func NewWorkshop(courseName string, instructor Instructor) Workshop {
	w := Workshop{}
	w.Course.Name = courseName
	w.Course.Slug = "workshop-" + courseName
	w.Instructor = instructor
	return w
}
