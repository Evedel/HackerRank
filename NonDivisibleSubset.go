package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the sockMerchant function below.
// Given a set of distinct integers,
// print the size of a maximal subset of S'
// where the sum of any 2 numbers in S' is not evenly divisible by k.
// 
// For example, the array S=[19,10,12,10,24,25,22] and k=4.
// One of the arrays that can be created is S'=[10,15,25].
// Another is S''=[19,22,24]. After testing all permutations,
// the maximum length solution array has 3 elements.
// 
// https://www.hackerrank.com/challenges/non-divisible-subset/problem
//
//  O(n)
//
func sockMerchant(n int, k int, arr []int32) int {
    hasZeroRems := false
    hasHalfRems := false
    rems := []int{}
    // (a + b)%k == 0 --> a%k + b%k = c%k == 0 || a%k + b%k = k

    // arr [1,7,4,2] to reminders only [1,1,1,2]
    for _, a := range(arr) {
        rems = append(rems, int(a)%k)
        if (rems[len(rems)-1] ==  0) {hasZeroRems = true}
        if (k%2 == 0) && (rems[len(rems)-1] == int(k/2)) {hasHalfRems = true}
    }

    // [1,1,1,2] need to match pairs to; i.e. 1+2 = 4%3+2%3 only one can be left
    gn := 0
    for rk:= 1; float32(rk) < float32(k)/2; rk++ {
        lk := k - rk
        nrk := 0
        nlk := 0
        for i := range(rems) {
            if (rems[i] == rk) {nrk += 1}
            if (rems[i] == lk) {nlk += 1}
        }
        if (nrk > nlk) {
            gn += nrk
        } else {
            gn += nlk
        }
    }
    // [3, 3, 3, 3] && k = 3   --> rems=[0,0,0,0] ---> only one can be left
    // [2, 6, 10, 14] && k = 4 --> rems=[2,2,2,2] ---> only one can be left
    if (hasZeroRems) {gn += 1}
    if (hasHalfRems) {gn += 1}
    return gn
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    inputTemp := strings.Split(readLine(reader), " ")
    nTemp, err := strconv.ParseInt(inputTemp[0], 10, 64)
    checkError(err)
    n := int(nTemp)

    kTemp, err := strconv.ParseInt(inputTemp[1], 10, 64)
    checkError(err)
    k := int(kTemp)

    arTemp := strings.Split(readLine(reader), " ")

    var ar []int32
    for i := 0; i < int(n); i++ {
        arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
        checkError(err)
        arItem := int32(arItemTemp)
        ar = append(ar, arItem)
    }

    result := sockMerchant(n, k, ar)

    fmt.Fprintf(writer, "%d\n", result)

    writer.Flush()
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
