#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the biggerIsGreater function below.
# Lexicographical order is often known as alphabetical order
# when dealing with strings.
# A string is greater than another string
# if it comes later in a lexicographically sorted list.
# 
# Given a word, create a new word by swapping some or all of its characters.
# This new word must meet two criteria:
# 
#    - It must be greater than the original word
#    - It must be the smallest word that meets the first condition
# For example, given the word abcd, the next largest word is abdc.
# And for the dkhc it is hcdk.
# 
# Complete the function biggerIsGreater below to create and return the new string
# meeting the criteria. If it is not possible, return 'no answer'.
# 
# https://www.hackerrank.com/challenges/bigger-is-greater
# 
# O(n+n*log(n))
# 
def biggerIsGreater(w):

    if (len(w) == 1): return 'no answer'

    if (len(w) == 2):
        if w[0] >= w[1]:
            return 'no answer'
        else:
            return w[1]+w[0]

    swapped = False
    for i in range(len(w)-2,-1,-1):
        minw = w[i+1]
        minj = i+1
        for j in range(i+1,len(w)):
            if (w[i] < w[j]) and (w[j] < minw):
                minw = w[j]
                minj = j
        if (minw > w[i]):
            w = w[:i]+w[minj]+w[i+1:minj]+w[i]+w[minj+1:]
            swapped = True
            break

    if swapped:
        w = w[:i+1] + ''.join(sorted(w[i+1:]))
        return w
    else:
        return 'no answer'

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    T = int(input())

    for T_itr in range(T):
        w = input()

        result = biggerIsGreater(w)

        fptr.write(result + '\n')

    fptr.close()
