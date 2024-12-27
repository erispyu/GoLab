package main

import "fmt"

type Route struct {
	Id       uint64
	Priority uint64
	Active   bool
}

func main() {
	routeList := []Route{
		{Id: 1, Priority: 1},
		{Id: 2, Priority: 2, Active: true},
		{Id: 3, Priority: 3},
	}
	// for _, r := range routeList {
	// 	if r.Id == 3 {
	// 		r.Active = true
	// 	}
	// }

	for i, r := range routeList {
		if r.Id == 3 {
			routeList[i].Active = true
		}
	}

	fmt.Println(routeList[2].Active)
}
