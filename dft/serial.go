package dft

import (
	"fmt"
	"math"
	"time"
)

func serial_calculate_dft(points []float64, N int, length int) []complexNum {
	var resDFT []complexNum = make([]complexNum, N)

	for k := 0; k < N; k++ {
		for n := 0; n < length; n++ {
			var theta float64 = (2 * math.Pi * float64(k*n)) / float64(N)
			resDFT[k].real += points[n] * math.Cos(theta)
			resDFT[k].img = resDFT[k].img - (points[n] * math.Sin(theta))
		}
	}
	return resDFT
}

func serial_calculate_idft(points []complexNum, N int, length int) []float64 {
	var resIDFT []float64 = make([]float64, length)

	for n := 0; n < N; n++ {
		resIDFT[n] = 0
		for k := 0; k < N; k++ {
			var theta float64 = (2 * math.Pi * float64(k*n)) / float64(N)
			resIDFT[n] = resIDFT[n] + points[k].real*math.Cos(theta) + points[k].img*math.Sin(theta)
		}
		resIDFT[n] = resIDFT[n] / float64(N)
	}
	return resIDFT
}

func Serial_Implementation(n int) {
	var discretePoints []float64 = make([]float64, n)
	var postDFTPoints []complexNum
	var postIDFTPoints []float64

	theta := 2 * math.Pi
	for i := 0; i < n; i++ {
		discretePoints[i] = (math.Sin((theta * float64(i)) / float64(n)))
	}

	writeToFile("data.txt", discretePoints)

	start := time.Now()

	postDFTPoints = serial_calculate_dft(discretePoints, n, n)

	postIDFTPoints = serial_calculate_idft(postDFTPoints, n, n)

	timeElapsed := time.Since(start).Seconds() * 1000

	fmt.Println(timeElapsed)

	writeToFile("dataIDFT.txt", postIDFTPoints)
}
