package main

type Num struct {
	Value int
}

type NumSliceDecrease []Num

func (s NumSliceDecrease) Len() int {
	return len(s)
}

func (s NumSliceDecrease) Less(i, j int) bool {
	return s[i].Value > s[j].Value
}

func (s NumSliceDecrease) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
