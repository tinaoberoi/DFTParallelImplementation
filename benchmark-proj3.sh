#!/bin/bash
#
#SBATCH --mail-user=toberoi@uchicago.edu
#SBATCH --mail-type=ALL
#SBATCH --job-name=project2-time
#SBATCH --output=/home/toberoi/project-3-tinaoberoi/proj3/time.txt
#SBATCH --error=./%j.%N.stderr
#SBATCH --chdir=/home/toberoi/project-3-tinaoberoi/proj3
#SBATCH --partition=debug
#SBATCH --nodes=1
#SBATCH --ntasks=1
#SBATCH --cpus-per-task=16
#SBATCH --mem-per-cpu=900
#SBATCH --exclusive
#SBATCH --time=03:59:00


go run editor.go 1000 s

echo ""

python3 plot.py 1000 s

go run editor.go 1000 p balance 2
go run editor.go 1000 p balance 4
go run editor.go 1000 p balance 6
go run editor.go 1000 p balance 8
go run editor.go 1000 p balance 12
go run editor.go 1000 p balance 16
go run editor.go 1000 p balance 32

echo ""

python3 plot.py 1000 pb

go run editor.go 1000 p steal 2
go run editor.go 1000 p steal 4
go run editor.go 1000 p steal 6
go run editor.go 1000 p steal 8
go run editor.go 1000 p steal 12
go run editor.go 1000 p steal 16
go run editor.go 1000 p steal 32

echo ""

python3 plot.py 1000 ps

echo "-------------------"

go run editor.go 10000 s

echo ""

python3 plot.py 10000 s

go run editor.go 10000 p balance 2
go run editor.go 10000 p balance 4
go run editor.go 10000 p balance 6
go run editor.go 10000 p balance 8
go run editor.go 10000 p balance 12
go run editor.go 10000 p balance 16
go run editor.go 10000 p balance 32

echo ""

python3 plot.py 10000 pb

go run editor.go 10000 p steal 2
go run editor.go 10000 p steal 4
go run editor.go 10000 p steal 6
go run editor.go 10000 p steal 8
go run editor.go 10000 p steal 12
go run editor.go 10000 p steal 16
go run editor.go 10000 p steal 32

echo ""

python3 plot.py 10000 ps

echo "-------------------"

go run editor.go 100000 s

echo ""

python3 plot.py 100000 s

go run editor.go 100000 p balance 2
go run editor.go 100000 p balance 4
go run editor.go 100000 p balance 6
go run editor.go 100000 p balance 8
go run editor.go 100000 p balance 12
go run editor.go 100000 p balance 16
go run editor.go 100000 p balance 32

echo ""

python3 plot.py 100000 pb

go run editor.go 100000 p steal 2
go run editor.go 100000 p steal 4
go run editor.go 100000 p steal 6
go run editor.go 100000 p steal 8
go run editor.go 100000 p steal 12
go run editor.go 100000 p steal 16
go run editor.go 100000 p steal 32

echo ""

python3 plot.py 10000 ps

python3 plot_speedup.py