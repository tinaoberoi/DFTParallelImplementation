import numpy as np
import matplotlib.pyplot as plt
import sys

a = np.array([], dtype=np.float64)
b = np.array([], dtype=np.float64)

n = sys.argv[1]
mode = sys.argv[2]
# print(f"N = {n}")

with open('data.txt') as f:
    lines = [float(line.rstrip()) for line in f]
    for line in lines :
        a = np.append(a, np.float32(line))
    # print(lines)
    
with open('dataIDFT.txt') as f:
    lines = [float(line.rstrip()) for line in f]
    for line in lines :
        x = np.float64(line)
        # print(f"line {line} and x {x}")
        b = np.append(b, x)
    # print(lines)
    
# print(a)
# print(b)

x = np.array([])
for i in range(int(n)):
    x = np.append(x, i)
    
# Plotting the Graph
plt.plot(x, a, color='green', label='original')
plt.plot(x, b, color='red', linestyle='dashed', label='IDFT')
plt.title("Curve plotted using the given points")
plt.xlabel("X")
plt.ylabel("Y")
plt.legend()
plt.savefig('dft_plot_' + str(n) + '_' + mode +  '.png')