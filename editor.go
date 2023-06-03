package main

import (
	"fmt"
	"os"
	"proj3/dft"
	"strconv"
)

const usage = "editor mode [seqLen] [mode] [algo] [number of threads] \n" +
	"seqLen	= Number of data points used for DFT - IDFT calculation.\n" +
	"mode     = (s) Serial and (p) Parallel \n" +
	"algo     = (balance) Call thread balancing exec and (steal) Call thread stealing executor \n" +
	"[number of threads] = Runs the parallel version of the program with the specified number of threads.\n"

func main() {
	if len(os.Args) < 2 {
		fmt.Println(usage)
		return
	}
	n, _ := strconv.Atoi(os.Args[1])
	mode := os.Args[2]
	if mode == "s" {
		// serial implementation
		dft.Serial_Implementation(n)

	} else {
		// parallel implementation
		algo := os.Args[3]
		num_threads, _ := strconv.Atoi(os.Args[4])
		if algo == "steal" {
			dft.StealExecutorService(num_threads, n)
		} else {
			dft.BalanceExecutorService(num_threads, n)
		}
	}
}
