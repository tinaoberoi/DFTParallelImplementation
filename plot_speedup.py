import numpy as np
import matplotlib.pyplot as plt
import sys


def plot_data(times):
    data_balanced = times[1: 8]
    data_steal = times[8:15]
    # print("Balanced ", data_balanced)
    # print("Steal ", data_steal)
    # print(times[0])
    serial_time = float(times[0])

    y_balanced = [1]
    y_steal = [1]
    
    for i, elem in enumerate(data_balanced):
        y_balanced.append(float(serial_time)/float(elem))
            
    
    for i, elem in enumerate(data_steal):
            y_steal.append(float(serial_time)/float(elem))
    
    # print("speedup balanced : ", y_balanced)
    # print("speedup steal : ", y_steal)
    return y_balanced, y_steal   

times = t = [ [0]*15 for i in range(4)]

with open('time.txt') as f:
    lines = f.read()
    lines = lines.split("-------------------")
    for i, line in enumerate(lines):
        sets = line.split()
        times[i] = sets
        # print(sets)
        
x = [1, 2, 4, 6, 8, 12, 16, 32]

#for n = 1000
data_0 = times[0]
y_b_0, y_s_0 = plot_data(data_0)

plt.plot(x, y_b_0, color='blue', label='balanced', marker='*')
plt.plot(x, y_s_0, color='red', linestyle='dashed', label='steal', marker='*')
plt.title("Speedup Graph N = 1000")
plt.xlabel("num_threads")
plt.ylabel("Speedup")
plt.legend()
plt.savefig('speedup_plot_' + str(1000) + '.png') 

plt.figure().clear()
plt.close()
plt.cla()
plt.clf()
    
# for n = 10000
data_1 = times[1]
y_b_1, y_s_1 = plot_data(data_1)

plt.plot(x,  y_b_1, color='blue', label='balanced', marker='*')
plt.plot(x, y_s_1, color='red', linestyle='dashed', label='steal', marker='*')
plt.title("Speedup Graph N = 10000")
plt.xlabel("num_threads")
plt.ylabel("Speedup")
plt.legend()
plt.savefig('speedup_plot_' + str(10000) + '.png') 

plt.figure().clear()
plt.close()
plt.cla()
plt.clf()

# for n = 100000
data_2 = times[2]
y_b_2, y_s_2 = plot_data(data_2)

plt.plot(x,  y_b_2, color='blue', label='balanced', marker='*')
plt.plot(x, y_s_2, color='red', linestyle='dashed', label='steal', marker='*')
plt.title("Speedup Graph N = 100000")
plt.xlabel("num_threads")
plt.ylabel("Speedup")
plt.legend()
plt.savefig('speedup_plot_' + str(100000) + '.png') 

plt.figure().clear()
plt.close()
plt.cla()
plt.clf() 