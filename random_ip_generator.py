#!/usr/bin/python3

from random import randint
import sys


if len(sys.argv) < 2:
    print("Usage : {} n (n : number of IPs to generate)".format(sys.argv[0]))
    exit(-1)

n = int(sys.argv[1])
for i in range(n):
    print("{}.{}.{}.{}".format(randint(0, 255), randint(0, 255), randint(0, 255), randint(0, 255)))