#!/usr/bin/env python3

import sys
import matplotlib
import matplotlib.pyplot as plt

font = {'family' : 'normal',
        'weight' : 'bold',
                'size'   : 12}

matplotlib.rc('font', **font)

matplotlib.use('svg')
plt.rcParams["font.family"] = "Times New Roman"

sizes      = [5, 10, 50, 100,  500,  1000,   2500]
times_arit = [0,  1, 13,  47, 3143, 18556, 238615]
times_dict = [1,  1,  1,   1,    1,     1,      1]

plt.title("Comparaisons des Temps d'Exécution")
plt.xlabel("Taille des données (octet)")
plt.ylabel("Temps d'exécution (us)")
plt.plot(sizes, times_arit, label="Encodeur Arithmétique")
plt.plot(sizes, times_dict, label="Encodeur LZW")
plt.legend()
plt.plot()
plt.savefig(sys.stdout.buffer)
