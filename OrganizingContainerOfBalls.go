package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)
func str(i int) string { return strconv.Itoa(i) }

func afterSwapPrint(c [][]int, n,i,j,k,l int) {
    fmt.Print(str(l) + ": C["+str(j)+"]["+str(i)+
        "] -> C["+str(i)+"]["+str(i)+"] and C["+str(i)+
        "]["+str(k)+"] -> C["+str(j)+"]["+str(k)+"]\n")

    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            fmt.Printf("%12d", c[i][j])
        }
        fmt.Print("\n")
    }
}

// Complete the organizingContainers function below.
// David has several containers, each with a number of balls in it.
// He has just enough containers to sort each type of ball he has into its own container.
// David wants to sort the balls using his sort method.
// 
// As an example, David has N containers and N different types of balls,
// both of which are numbered from 1 to N.
// The distribution of ball types per container are described by an NxN matrix of integers.
// 
// In a single operation, David can swap two balls located in different containers.
// 
// David wants to perform some number of swap operations such that:
//      - Each container contains only balls of the same type.
//      - No two balls of the same type are located in different containers.
// You must perform N queries where each query is in the form of a matrix Q.
// For each query, print "Possible" on a new line
// if David can satisfy the conditions above for the given matrix.
// Otherwise, print "Impossible".
// 
// https://www.hackerrank.com/challenges/organizing-containers-of-balls
// 

func organizingContainers(container [][]int) string {
    n := len(container)
    c := container
    totBalls := []int{}
    totPlace := []int{}
    for i := 0; i < n; i++ {
        totPlace = append(totPlace,0)
        for j := 0; j < n; j++ {
            if (i == 0) {totBalls = append(totBalls,0)}
            totPlace[i] += c[i][j]
            totBalls[j] += c[i][j]
        }
    }
    for i := 0; i < n; i++ {
        foundPlace := false
        for j := 0; j < n; j++ {
            if (totBalls[i] == totPlace[j]) {
                totPlace[j] = -1
                foundPlace = true
                break
            }
        }
        if (!foundPlace) {return "Impossible"}
    }
    return "Possible"
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    q := int32(qTemp)

    for qItr := 0; qItr < int(q); qItr++ {
        nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
        checkError(err)
        n := int32(nTemp)

        var container [][]int
        for i := 0; i < int(n); i++ {
            containerRowTemp := strings.Split(readLine(reader), " ")

            var containerRow []int
            for _, containerRowItem := range containerRowTemp {
                containerItemTemp, err := strconv.ParseInt(containerRowItem, 10, 64)
                checkError(err)
                containerItem := int(containerItemTemp)
                containerRow = append(containerRow, containerItem)
            }

            if len(containerRow) != int(int(n)) {
                panic("Bad input")
            }

            container = append(container, containerRow)
        }

        result := organizingContainers(container)

        fmt.Fprintf(writer, "%s\n", result)
    }

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
