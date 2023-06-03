# Project 3

Write up : Project3WriteUp.pdf

Peanut Script: benchmark-proj3.sh

Code for dft implementation: dft/*

Interface: editor.go

```
Usage: editor mode [seqLen] [mode] [algo] [number of threads] 
seqLen  = Number of data points used for DFT - IDFT calculation. 
mode     = (s) Serial and (p) Parallel  
algo     = (balance) Call thread balancing exec and (steal) Call thread stealing executor  
number of threads = Runs the parallel version of the program with the specified number of threads. 
```
Images: ./images/*

Scripts: 
- plot.py (plots correctness graph)
- plot_speedup.py (plots speedup graphs from time.txt from N = 1000, 10000, 100000)

