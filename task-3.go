package main

// import (
// 	"fmt"
// 	"math"
// )

// type hitung2d interface {
// 	luas() float64
// 	keliling() float64
// }
// type hitung3d interface {
// 	volume() float64
// }
// type hitung interface {
// 	hitung2d
// 	hitung3d
// }
// type kerucut struct {
// 	tinggi, phi, r float64
// }

// func (k kerucut) luas() float64 {
// 	return k.tinggi * (2 * k.r)
// }

// func (k kerucut) keliling() float64 {
// 	return math.Sqrt(k.tinggi+k.r)*2 + k.r*2
// }

// func (k kerucut) volume() float64 {
// 	return k.phi * k.r * k.r * k.tinggi / 3
// }

// func hitungKerucut() {
// 	k := kerucut{tinggi: 11, r: 7, phi: 3.14}

// 	fmt.Println("luas segitiga =", k.luas())
// 	fmt.Println("keliling segitiga =", k.keliling())
// 	fmt.Println("volume segitiga =", k.volume())
// }
