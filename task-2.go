package main

import "fmt"

func deretBilangan(num int) []int {
	deretAngka := []int{}
	for i := 1; i <= 40; i++ {
		deretAngka = append(deretAngka, i)
	}

	return deretAngka
}

type derett struct {
	data []int
}

func (d derett) prima() {
	var y []int = []int{}
	devider := 0
	for _, v := range d.data {
		devider = 0
		for i := 1; i <= v; i++ {
			if v%i == 0 {
				devider++
			}
		}
		if devider == 2 {
			y = append(y, v)
		}
	}
	fmt.Println(y)
}

func (d derett) ganjil() {
	x := 0
	var y []int = []int{}
	for _, v := range d.data {
		x = v + v - 1
		if x < len(d.data) {
			y = append(y, x)
		} else {
			break
		}
	}
	fmt.Println(y)
}

func (d derett) genap() {
	var y []int = []int{}
	for _, v := range d.data {
		if v%2 == 0 {
			y = append(y, v)
		}
	}
	fmt.Println(y)
}

func (d derett) fibonacci() {
	x := 0
	var y []int = []int{}
	for i := 0; i < len(d.data); i++ {
		x = 0
		if i <= 1 {
			y = append(y, d.data[i])
		} else if i > 0 {
			x = y[i-1] + y[i-2]
			if x > len(d.data) {
				break
			}
			y = append(y, x)
		}
	}

	fmt.Println(y)
}
