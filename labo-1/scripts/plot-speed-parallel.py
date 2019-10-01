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

workers = [     1,     4,     8,    16,    32,    64]
times   = [147106, 48962, 26621, 18218, 18651, 20175]

plt.title("Parallélisation de l'Encodeur Arithmétique")
plt.xlabel("Nombre de fils d'exécutions")
plt.ylabel("Temps d'exécution (us)")
plt.plot(workers, times)
plt.plot()
plt.savefig(sys.stdout.buffer)
