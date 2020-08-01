# you can write to stdout for debugging purposes, e.g.
# print("this is a debug message")
# T is the array of tasks ["test1", "test2","test4a","test4b","test4c","test3"]
# R is the array of results ["OK","OK","OK","NOT OK","OK","NOT OK"]
# The task is ok if all the subtasks are ok
# Find the total percentage of OKs
# 50/100 in this case
def solution(T, R):
    pref = 0
    while not T[0][pref].isdigit():
        pref += 1

    ntask = 0
    nok   = 0
    groups = {}
    for i in range(len(T)):
        if T[i][pref:].isdigit():
            ntask += 1
            if R[i] == "OK": nok   += 1
        else:
            k = 1
            while T[i][pref+k].isdigit(): k += 1
            if T[i][pref:pref+k] in groups:
                groups[T[i][pref:pref+k]] = groups[T[i][pref:pref+k]] and (R[i] == "OK")
            else:
                groups[T[i][pref:pref+k]] = (R[i] == "OK")
    
    for g in groups:
        ntask += 1
        if groups[g]: nok += 1
        
    return 0 if ntask == 0 else int(nok*100/ntask)