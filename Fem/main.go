package main

import (
	"fm.com/data"
	"fmt"
	"time"
)

func main() {
	max := data.Instructor{Id: 1, FirstName: "John", LastName: "Doe", Score: 98.5}
	max.FirstName = "Max"
	kyle := data.NewInstructor(2, "Kyle", "Smith", 95.5)
	fmt.Println(max.FullName())
	fmt.Println(kyle.FullName())

	goCourse := data.Course{Id: 1, Name: "Go Fundamentals", Slug: "go-fundamentals", Legacy: false, Duration: 3.5}
	fmt.Println(goCourse)

	swiftWs := data.Workshop{Course: goCourse, Date: time.Now()}

	fmt.Println(swiftWs)

	pythonWs := data.NewWorkshop("Python Fundamentals", kyle)

	fmt.Println(pythonWs.Course)

	courses := [2]data.KindOfSignable{goCourse, pythonWs}

	for _, c := range courses {
		fmt.Println(c)
	}
}
