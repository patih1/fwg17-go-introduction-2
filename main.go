package main

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
		deret,
	}
	data.prima()
	data.ganjil()
	data.genap()
	data.fibonacci()
	// data.cek()

	hitungKerucut()
}
