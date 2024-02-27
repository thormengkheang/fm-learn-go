package data

type Instructor struct {
	Id        int
	FirstName string
	LastName  string
	Score     float64
}

func NewInstructor(id int, firstName string, lastName string, score float64) Instructor {
	return Instructor{Id: id, FirstName: firstName, LastName: lastName, Score: score}
}

func (i Instructor) FullName() string {
	return i.FirstName + " " + i.LastName
}
