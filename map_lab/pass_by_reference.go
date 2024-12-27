package main

func TestMapReference(m map[string]int) {
	for k, v := range m {
		m[k] = v + 1
	}
}
