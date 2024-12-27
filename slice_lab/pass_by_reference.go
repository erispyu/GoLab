package main

func TestSliceReference(s []int) {
	for i, v := range s {
		s[i] = v + 1
	}
}
