#!/bin/python3

import math
import os
import random
import re
import sys

def checkMask(G,P,i,j,dxP,dyP):
    for ip in range(dxP):
        for jp in range(dyP):
            if P[ip][jp] != G[i+ip][j+jp]: return False
    return True

# Complete the gridSearch function below.
# Given a 2D array of digits or grid, try to find the occurrence of a given 2D pattern of digits.
def gridSearch(G, P):
    
    dxG = len(G)
    dyG = len(G[0])
    dxP = len(P)
    dyP = len(P[0])

    for i in range(dxG):
        for j in range(dyG):
            if G[i][j] == P[0][0]:
                if (i+dxP-1 < dxG) and (j+dyP-1 < dyG):
                    if checkMask(G,P,i,j,dxP,dyP): return "YES"
    return "NO"

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    t = int(input())

    for t_itr in range(t):
        RC = input().split()

        R = int(RC[0])

        C = int(RC[1])

        G = []

        for _ in range(R):
            G_item = input()
            G.append(G_item)

        rc = input().split()

        r = int(rc[0])

        c = int(rc[1])

        P = []

        for _ in range(r):
            P_item = input()
            P.append(P_item)

        result = gridSearch(G, P)

        fptr.write(result + '\n')

    fptr.close()
