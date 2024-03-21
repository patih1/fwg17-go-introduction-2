package main

import "fmt"

func main() {
	// task-1
	pembulatan(3.37)
	pembulatan(3.32)
	pembulatan(3.324)

	//==========================================//

	// task-2
	deret := deretBilangan(40)
	// fmt.Println(deret)
	var data deretAngka = deretAngka{
		&deret,
	}

	fmt.Println(data.calculateAll())

	// hitungKerucut()
}
