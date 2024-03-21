package main

import (
	"fmt"
	"strconv"
	"strings"
)

func pembulatan(num float64) {
	// floatStr := strconv.Itoa(num)
	floatStr := fmt.Sprintf("%v", num)
	strLastDec := "0.0"
	for i := 3; i < len(floatStr); i++ {
		strLastDec += strings.Split(floatStr, "")[i]
	}
	lastDec, _ := strconv.ParseFloat(strLastDec, 32)
	x := 0.1 - lastDec

	finalRes := ""

	if x <= 0.05 {
		result := fmt.Sprintf("%v", num+0.1)

		for i := 0; i < 3; i++ {
			finalRes += strings.Split(result, "")[i]
		}

		output, _ := strconv.ParseFloat(finalRes, 64)
		fmt.Println(output)
	} else {
		result := fmt.Sprintf("%v", num)

		for i := 0; i < 3; i++ {
			finalRes += strings.Split(result, "")[i]
		}

		output, _ := strconv.ParseFloat(finalRes, 64)
		fmt.Println(output)
	}
}
