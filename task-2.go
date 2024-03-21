package main

import "fmt"

func deretBilangan(num int) []int {
	deretAngka := []int{}
	for i := 0; i <= num; i++ {
		deretAngka = append(deretAngka, i)
	}

	return deretAngka
}

type deretAngka struct {
	data *[]int
}

func (d deretAngka) prima() []int {
	var results []int = []int{}
	devider := 0
	for _, v := range *d.data {
		devider = 0
		for i := 1; i <= v; i++ {
			if v%i == 0 {
				devider++
			}
		}
		if devider == 2 {
			results = append(results, v)
		}
	}
	return results
}

func (d deretAngka) ganjil() []int {
	var results []int = []int{}
	for _, v := range *d.data {
		if v%2 != 0 {
			if v > 0 {
				results = append(results, v)
			}
		}
	}
	return results
}

func (d deretAngka) genap() []int {
	var results []int = []int{}
	for _, v := range *d.data {
		if v%2 == 0 {
			if v > 0 {
				results = append(results, v)
			}
		}
	}
	return results
}

// [0, 1, 2, 3, 4, .... , 40]

func (d deretAngka) fibonacci() []int {
	nextPatern := 0
	var results []int = []int{}         // [0, 1, 1, 2, 3, 5, 8, 13]
	for i := 0; i < len(*d.data); i++ { // i = 3
		x := *d.data
		if i <= 1 {
			results = append(results, x[i])
		} else if i > 1 {

			nextPatern = results[i-1] + results[i-2]

			if nextPatern > len(*d.data)-1 {
				break
			}
			results = append(results, nextPatern)
		}
	}

	return results
}

func (d deretAngka) calculateAll() (string, string, string, string) {
	a := fmt.Sprint("prima: ", d.prima(), "\n")
	b := fmt.Sprint("ganjil: ", d.ganjil(), "\n")
	c := fmt.Sprint("genap: ", d.genap(), "\n")
	e := fmt.Sprint("fibonacci: ", d.fibonacci())

	return a, b, c, e
}
