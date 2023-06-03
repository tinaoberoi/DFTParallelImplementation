package dft

import (
	"fmt"
	"math"
	"proj3/concurrent"
	"time"
)

type goContextDFT struct {
	discrete_points []float64
	dft_points      []complexNum
	length          int
}

type goContextIDFT struct {
	dft_points  []complexNum
	idft_points []float64
	length      int
}

// Function to calculate dft
func calculate_dft(points []float64, dft_points []complexNum, k int, N int) {
	for n := 0; n < N; n++ {
		var theta float64 = (2 * math.Pi * float64(k*n)) / float64(N)
		dft_points[k].real += points[n] * math.Cos(theta)
		dft_points[k].img = dft_points[k].img - (points[n] * math.Sin(theta))
	}
}

// Function to calculate idft
func calculate_idft(points []complexNum, res []float64, n int, N int) {
	for k := 0; k < N; k++ {
		var theta float64 = (2 * math.Pi * float64(k*n)) / float64(N)
		res[n] += points[k].real*math.Cos(theta) + points[k].img*math.Sin(theta)
	}
	res[n] = res[n] / float64(N)
}

type DFTTask struct {
	ctx  *goContextDFT
	rank int
}

type IDFTTask struct {
	ctx  *goContextIDFT
	rank int
}

func NewDFTTask(ctx *goContextDFT, rank int) concurrent.Runnable {
	return &DFTTask{ctx, rank}
}

func NewIDFTTask(ctx *goContextIDFT, rank int) concurrent.Runnable {
	return &IDFTTask{ctx, rank}
}

// DFT run function
func (task *DFTTask) Run() {
	calculate_dft(task.ctx.discrete_points, task.ctx.dft_points, task.rank, task.ctx.length)
}

// IDFT run function
func (task *IDFTTask) Run() {
	calculate_idft(task.ctx.dft_points, task.ctx.idft_points, task.rank, task.ctx.length)
}

// Calls stealing algorithm
func StealExecutorService(threadCount int, N int) {
	executorDFT := concurrent.NewWorkStealingExecutor(threadCount, threadCount)
	executorIDFT := concurrent.NewWorkStealingExecutor(threadCount, threadCount)

	var discretePoints []float64 = make([]float64, N)
	var futures []concurrent.Future
	var context goContextDFT

	theta := 2 * math.Pi
	for i := 0; i < N; i++ {
		discretePoints[i] = (math.Sin((theta * float64(i)) / float64(N)))
	}

	writeToFile("data.txt", discretePoints)

	start := time.Now()

	context.discrete_points = discretePoints
	context.length = N
	context.dft_points = make([]complexNum, N)

	for i := 0; i < N; i++ {
		task := NewDFTTask(&context, i)
		futures = append(futures, executorDFT.Submit(task))

	}
	executorDFT.Shutdown()

	var futures2 []concurrent.Future
	var context2 goContextIDFT

	context2.length = N
	context2.dft_points = context.dft_points

	context2.idft_points = make([]float64, N)

	for i := 0; i < N; i++ {
		task := NewIDFTTask(&context2, i)
		futures2 = append(futures2, executorIDFT.Submit(task))

	}
	executorIDFT.Shutdown()

	timeElapsed := time.Since(start).Seconds() * 1000

	fmt.Println(timeElapsed)

	writeToFile("dataIDFT.txt", context2.idft_points)

}

// Calls balancing algorithm
func BalanceExecutorService(threadCount int, N int) {
	thresholdBalance := (N / threadCount) * 10
	executorDFT := concurrent.NewWorkBalancingExecutor(threadCount, threadCount, thresholdBalance)
	executorIDFT := concurrent.NewWorkBalancingExecutor(threadCount, threadCount, thresholdBalance)

	var discretePoints []float64 = make([]float64, N)
	var futures []concurrent.Future
	var context goContextDFT

	theta := 2 * math.Pi
	for i := 0; i < N; i++ {
		discretePoints[i] = (math.Sin((theta * float64(i)) / float64(N)))
	}

	writeToFile("data.txt", discretePoints)

	start := time.Now()

	context.discrete_points = discretePoints
	context.length = N
	context.dft_points = make([]complexNum, N)

	for i := 0; i < N; i++ {
		task := NewDFTTask(&context, i)
		futures = append(futures, executorDFT.Submit(task))

	}
	executorDFT.Shutdown()

	var futures2 []concurrent.Future
	var context2 goContextIDFT

	context2.length = N
	context2.dft_points = context.dft_points

	context2.idft_points = make([]float64, N)

	for i := 0; i < N; i++ {
		task := NewIDFTTask(&context2, i)
		futures2 = append(futures2, executorIDFT.Submit(task))

	}
	executorIDFT.Shutdown()

	timeElapsed := time.Since(start).Seconds() * 1000

	fmt.Println(timeElapsed)

	writeToFile("dataIDFT.txt", context2.idft_points)

}
