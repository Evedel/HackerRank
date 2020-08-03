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
func organizingContainers(container [][]int) string {
    n := len(container)
    c := container
    // afterSwapPrint(c,n,0,0,0,0)
    // for each basket id && target ball id
    for i := 0; i < n; i++ {
        // find out what is the number of wrong balls in this basket i
        nwrongballs := 0
        wrongballid := i+1
        for k := i+1; k < n; k++ {
            nwrongballs += c[i][k]
        }
        // visit every other basket j
        for j := i+1; j < n; j++ {
            // if this basket j contains balls i, need to swap them for any in basket i
            for c[j][i] != 0 {
                // But if there is nothing to swap with => "NO"
                if (nwrongballs == 0) {return "Impossible"}
                swapballs := 0
                if c[j][i] < nwrongballs {
                    swapballs = c[j][i]
                } else {
                    swapballs = nwrongballs
                }
                if (swapballs > c[i][wrongballid]) {
                    swapballs = c[i][wrongballid]
                }
                c[i][i] += swapballs
                c[j][i] -= swapballs
                c[j][wrongballid] += swapballs
                c[i][wrongballid] -= swapballs

                // afterSwapPrint(c,n,i,j,wrongballid,swapballs)
                
                nwrongballs -= swapballs
                if c[i][wrongballid] == 0 {wrongballid += 1}
            }
        }
        // if after all the basket j were visited,
        // we still have some wrong balls => "NO"
        if (nwrongballs != 0) {return "Impossible"}
    }
    // If no mistakes => "YES"
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
