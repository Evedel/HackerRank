#!/bin/python3

import math
import os
import random
import re
import sys

def binSearch(ranks, score):
    # print(ranks)

    maxS = ranks[1]
    maxI = 1
    minS = ranks[len(ranks)]
    minI = len(ranks)
    resI = (minI+maxI)//2 
    resS = ranks[resI]

    if (score > maxS): return maxI
    if (score < minS): return minI+1

    i = 0
    while ((minI-maxI) > 1) and (i < 10):
        # print(maxS, resS, minS, score)
        if (resS > score):
            maxI = resI
        else:
            minI = resI
        maxS = ranks[maxI]
        minS = ranks[minI]
        resI = (minI+maxI)//2 
        resS = ranks[resI]
        i += 1

    if (score >= maxS):
        return maxI
    elif (score >= minS):
        return minI

# Complete the climbingLeaderboard function below.
# Alice is playing an arcade game. 
# She wants to climb to the top of the leaderboard and
# wants to track her ranking.
# The game uses Dense Ranking, so its leaderboard works like this:
# 
# - The player with the highest score is ranked number 1 on the leaderboard.
# - Players who have equal scores receive the same ranking number,
#   and the next player(s) receive the immediately following ranking number.
# 
# For example, the four players on the leaderboard have high scores of 100, 90, 90, and 80.
# Those players will have ranks 1, 2, 2, and 3, respectively.
# If Alice's scores are 70, 80 and 105, her rankings after each game are 4, 3 and 1.
#
# O(n+m*log(n))
#
def climbingLeaderboard(scores, alice):
    ranks = {}
    irank = 1
    ranks[1] = scores[0]
    for i in range(1,len(scores)):
        if scores[i] != ranks[irank]:
            irank += 1
            ranks[irank] = scores[i]
    
    result = []
    for r in alice:
        result.append(binSearch(ranks,r))
    
    # print(ranks)
    return result

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    scores_count = int(input())

    scores = list(map(int, input().rstrip().split()))

    alice_count = int(input())

    alice = list(map(int, input().rstrip().split()))

    result = climbingLeaderboard(scores, alice)

    fptr.write('\n'.join(map(str, result)))
    fptr.write('\n')

    fptr.close()
