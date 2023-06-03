package dft

import (
	"fmt"
	"os"
)

type complexNum struct {
	real, img float64
}

func (num *complexNum) printNum() {
	fmt.Println(num.real, " + ", num.img, "j")
}

func writeToFile(filename string, data []float64) {
	file, errs := os.Create(filename)
	if errs != nil {
		fmt.Println("Failed to create file:", errs)
		return
	}
	defer file.Close()

	for _, elem := range data {
		s := fmt.Sprintln(elem)
		_, err := file.WriteString(s)
		if err != nil {
			fmt.Printf("error writing string: %v", err)
		}
	}

}
