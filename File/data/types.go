package data

import "fmt"

type integer = int                 // type alias
type json = map[string]interface{} //  type alias

var integor integer = 10

// not a type alias it create new type distance
type distance float64   // in miles
type distanceKM float64 // in kilometers

// method attached to distance type
func (miles distance) toKM() distanceKM {
	return distanceKM(miles * 1.60934)
}

func (km distanceKM) toMiles() distance {
	return distance(km / 1.60934)
}

func test() {
	d := distance(10.0)
	km := d.toKM()
	miles := km.toMiles()
	fmt.Println(miles)
}
