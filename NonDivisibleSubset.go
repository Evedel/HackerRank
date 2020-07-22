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
func sockMerchant(n int32, k int32, arr []int32) int32 {
    tmp   := arr[:]
    pairs := []int{}
    res   := 0
    for i := 0; i < len(tmp); i++ { pairs = append(pairs,0) }
    
    for len(tmp) > 0 {
        for i := 0; i < len(tmp); i++ { pairs[i] = 0 }

        for i := 0; i < len(tmp); i++ {
            for j := i+1; j < len(tmp); j++ {
                if (tmp[i]+tmp[j])%k == 0 {
                    pairs[i] += 1
                    pairs[j] += 1
                }
            }
        }
        fmt.Print(pairs, "\n")
        maxI := -1
        maxV :=  0
        for i := 0; i < len(tmp); i++ {
            if pairs[i] > maxV {
                maxI = i
                maxV = pairs[i]
            }
        }
        if (maxI != -1) {
            tmp = append(tmp[:maxI], tmp[maxI+1:]...)
            pairs = append(pairs[:maxI], pairs[maxI+1:]...)
        } else {
            return int32(len(tmp) + res)
        }
        for i := len(tmp)-1; i >= 0; i-- {
            if pairs[i] == 0 {
                res += 1
                tmp = append(tmp[:i], tmp[i+1:]...)
                pairs = append(pairs[:i], pairs[i+1:]...)
            }
        }
    }
    return int32(0)
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
    n := int32(nTemp)

    kTemp, err := strconv.ParseInt(inputTemp[1], 10, 64)
    checkError(err)
    k := int32(kTemp)

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
