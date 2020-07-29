#!/bin/python3

import math
import os
import random
import re
import sys

# Given a grid of size NxM, each cell in the grid is either GOOD or BAD.
# 
# A valid plus is defined here as the crossing of two segments (horizontal and vertical)
# of equal lengths. These lengths must be odd, and the middle cell of its horizontal
# segment must cross the middle cell of its vertical segment.
# 
# Find the two largest valid pluses that can be drawn on GOOD cells in the grid,
# and return an integer denoting the maximum product of their areas.
# 
# Note: The two pluses cannot overlap, and the product of their areas should be maximal.
# 
# 
# https://www.hackerrank.com/challenges/two-pluses
# 

def getPlusSize(grid,i,j,lx,ly):
    for k in range(1,min(lx,ly)):
        if not (
            (i-k >= 0) and
            (j-k >= 0) and
            (i+k < lx) and
            (j+k < ly)
        ):
            return k-1
        if not (
            (grid[i-k][j] == 1) and
            (grid[i][j-k] == 1) and
            (grid[i+k][j] == 1) and
            (grid[i][j+k] == 1)
        ):
            return k-1

def makeUnavaliable(grid,i,j,size):
    for k in range(1,size+1):
        grid[i-k][j] = 0
        grid[i][j-k] = 0
        grid[i+k][j] = 0
        grid[i][j+k] = 0
    return grid

def makeAvaliable(grid,i,j,size):
    for k in range(1,size+1):
        grid[i-k][j] = 1
        grid[i][j-k] = 1
        grid[i+k][j] = 1
        grid[i][j+k] = 1
    return grid

def area(size):
    return 1+size*4

# Complete the twoPluses function below.
def twoPluses(grid):
    lx = len(grid)
    ly = len(grid[0])
    gr = []
    for i in range(lx):
        gr.append([])
        for j in range(ly):
            gr[i].append(1 if grid[i][j]=='G' else 0)

    maxArea = 0
    for i1 in range(lx):
        for j1 in range(ly):
            if (gr[i1][j1] == 1):
                size1 = getPlusSize(gr,i1,j1,lx,ly)
                for s in range(size1+1):
                    gr2 = makeUnavaliable(gr,i1,j1,s)
                    size2 = 0
                    for i2 in range(lx):
                        for j2 in range(ly):
                            if (gr2[i2][j2] == 1):
                                tmpSize = getPlusSize(gr2,i2,j2,lx,ly)
                                if (tmpSize > size2): size2 = tmpSize
                    tmpArea = area(s)*area(size2)
                    if (tmpArea > maxArea): maxArea = tmpArea
                    gr2 = makeAvaliable(gr,i1,j1,s)
    
    return maxArea

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    nm = input().split()

    n = int(nm[0])

    m = int(nm[1])

    grid = []

    for _ in range(n):
        grid_item = input()
        grid.append(grid_item)

    result = twoPluses(grid)

    fptr.write(str(result) + '\n')

    fptr.close()
