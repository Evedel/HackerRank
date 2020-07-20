ckage main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

func swapOrRev(a, b int) string {
	// prefer swap over reverse where possible
	// [2,1]     -> swap 1 2
	// [3,2,1]   -> swap 1 3
	// [4,3,2,1] -> reverse 1 4
    if a == b+1 {
        return "swap "+strconv.Itoa(b)+" "+strconv.Itoa(a)
    }
    if a == b+2 {
        return "swap "+strconv.Itoa(b)+" "+strconv.Itoa(a)
    }
    return "reverse "+strconv.Itoa(b)+" "+strconv.Itoa(a)
}

// Given an array of integers, determine whether the array can be sorted in ascending order
// using only one of the following operations one time.
// - Swap two elements.
// - Reverse one sub-segment.
// Determine whether one, both or neither of the operations will complete the task.
// If both work, choose swap.
// https://www.hackerrank.com/challenges/almost-sorted/problem
//
// O(n)
//

func almostSorted(arr []int32) {
    if (len(arr) < 2) {
        fmt.Print("yes\n")
        return
    }

    intervals := []int{}
    intervals = append(intervals, 0)
     // T if ascending [1,2,3], F if descending [3,2,1]
    isAscendingSeq := arr[0] < arr[1]
    i := 2
    for i < int(len(arr)) {
        if (arr[i] > arr[i-1]) != isAscendingSeq {
            intervals = append(intervals, i-1)
            isAscendingSeq = !isAscendingSeq
        }
        i += 1
    }
    
    // fmt.Print(intervals,"\n")

    if (len(intervals) == 1) {
		// Single interval. Either sorted or not.
        if isAscendingSeq {
            fmt.Print("yes")
        } else {
            fmt.Print("yes\n")
            fmt.Print(swapOrRev(len(arr),1))
        }
    } else if (len(intervals) == 2) {
		// Two intervals /\ or \/. Try to swap descending part in both cases.
        if isAscendingSeq {
            if arr[intervals[0]] > arr[intervals[1]+1] {
                fmt.Print("no\n")
            } else {
                fmt.Print("yes\n")
                fmt.Print(swapOrRev(intervals[1]+1, intervals[0]+1))
            }
        } else {
            if arr[intervals[1]-1] > arr[len(arr)-1] {
                fmt.Print("no\n")
            } else {
                fmt.Print("yes\n")
                fmt.Print(swapOrRev(len(arr), intervals[1]+1))
            }
        }
    } else if (len(intervals) == 3) {
		// Either /\/ or \/\.
		// The second case is only possible to fix in one operation,
		//     if it is a swap of the first and the last elements
		// In the first case, it may be possible to swap or reverse the middle section
        if isAscendingSeq {
            if (arr[intervals[1]-1] > arr[intervals[2]+1]) {
                fmt.Print("no\n")
            } else {
                if ((arr[intervals[1]-1] > arr[intervals[2]]) ||
                    (arr[intervals[1]] > arr[intervals[2]+1])) {
                    fmt.Print("no\n")
                } else {
                    fmt.Print("yes\n")
                    fmt.Print(swapOrRev(intervals[2]+1, intervals[1]+1))
                }
            }
        } else {
            if (intervals[1] == 1) && (intervals[2] == len(arr)-2) {
                if (arr[0] > arr[len(arr)-2]) &&
					(arr[len(arr)-1] < arr[1]) {
                    fmt.Print("yes\n")
                    fmt.Print("swap ", 1, len(arr))
                } else {
                    fmt.Print("no")
                }
            } else {
                fmt.Print("no")
            }
        }
    } else if (len(intervals) == 4) {
		//         /\/\ or \/\/
		// only possibel when it is a swap between the last (first) element and some random element
        if isAscendingSeq {
            el1  := arr[intervals[0]]
            el1r := arr[intervals[0]+1]
            el2  := arr[intervals[3]]
            el2l := arr[intervals[3]-1]
            el2r := arr[intervals[3]+1]
            if (el1 > el2) &&
                (el2 < el1r) &&
                (el1 > el2l) && (el1 < el2r) {
                    fmt.Print("yes\n")
                    fmt.Print("swap ", intervals[0]+1, intervals[3]+1)
            } else {
                fmt.Print("no")
            }
        } else {
            el1  := arr[intervals[1]]
            el1l := arr[intervals[1]-1]
            el1r := arr[intervals[1]+1]
            el2  := arr[len(arr)-1]
            el2l := arr[len(arr)-1-1]
            if (el1 > el2) &&
                (el2 > el1l) && (el2 < el1r) &&
                (el1 > el2l) {
                    fmt.Print("yes\n")
                    fmt.Print("swap ", intervals[1]+1, len(arr))
            } else {
                fmt.Print("no")
            }
        }
    } else if (len(intervals) == 5) {
		//       \/\/\  or  /\/\/
		// either it is a swap of two random elements or it is impossible to sort in one operation
        if isAscendingSeq {
            el1  := arr[intervals[1]]
            el1l := arr[intervals[1]-1]
            el1r := arr[intervals[1]+1]
            el2  := arr[intervals[4]] 
            el2l := arr[intervals[4]-1]
            el2r := arr[intervals[4]+1]
            if (el1 > el2) &&
                (el2 > el1l) && (el2 < el1r) &&
                (el1 > el2l) && (el1 < el2r) {
                    fmt.Print("yes\n")
                    fmt.Print("swap ", intervals[1]+1, intervals[4]+1)
            } else {
                fmt.Print("no")
            }
        } else {
            fmt.Print("no\n")
        }
    } else{
		// is not almost sorted
        fmt.Print("no\n")
    }
    return
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    arrTemp := strings.Split(readLine(reader), " ")

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    almostSorted(arr)
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

