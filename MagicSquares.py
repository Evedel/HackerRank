#!/bin/python3

import math
import os
import random
import re
import sys

def rotate(s):
    return [list(elem) for elem in zip(*s[::-1])]

def mirrorH(s):
    t    = s[0][:]
    s[0] = s[2][:]
    s[2] = t[:]
    return s

def mirrorV(s):
    s2 = [[],[],[]]
    for i in range(3):
        s2[i]   = s[i][:]
        t       = s2[i][0]
        s2[i][0] = s2[i][2]
        s2[i][2] = t
    return s2

def draw(s):
    t = ""
    for i in range(3):
        for j in range(3):
            t += str(s[i][j]) + " "
        print(t)
        t = ""

# Complete the formingMagicSquare function below.
# We define a magic square to be an 3x3 matrix of distinct positive integers
# from 1 to 9 where the sum of any row, column, or diagonal of length 
# is always equal to the same number: the magic constant.
# You will be given a 3x3 matrix S of integers in the inclusive range [1,9].
# We can convert any digit `a` to any other digit `b` at cost of |a-b|.
# Given S, convert it into a magic square at minimal cost.
# Print this cost on a new line.
def formingMagicSquare(s):
    magic = [[8,1,6], [3,5,7], [4,9,2]]
    
    universe = []
    universe.append(magic)

    rotated = rotate(magic)
    for i in range(3):
        universe.append(rotated)
        rotated = rotate(rotated)
    
    flipped = mirrorH(magic[:])
    universe.append(flipped)

    rotated = rotate(flipped)
    for i in range(3):
        universe.append(rotated)
        rotated = rotate(rotated)

    flipped = mirrorV(magic)
    universe.append(flipped)

    rotated = rotate(flipped)
    for i in range(3):
        universe.append(rotated)
        rotated = rotate(rotated)
    
    globDiff = 1000000
    for u in range(len(universe)):
        locDiff = 0
        for i in range(3):
            for j in range(3):
                    locDiff += abs(universe[u][i][j] - s[i][j])
        if (locDiff < globDiff):
            globDiff = locDiff

    return globDiff

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    s = []

    for _ in range(3):
        s.append(list(map(int, input().rstrip().split())))

    result = formingMagicSquare(s)

    fptr.write(str(result) + '\n')

    fptr.close()

