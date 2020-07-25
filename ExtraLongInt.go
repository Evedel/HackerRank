package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "math/big"
)

// Complete the extraLongFactorials function below.
// Calculate and print the factorial of a given integer.
// 1 <= n <= 100
// n = 45 
// factorial = 119622220865480194561963161495657715064383733760000000000
func extraLongFactorials(n int) {
    arr := []int{}
    arr = append(arr, 1)
    for i := 2; i <= n; i++ {
        rem := 0
        for j := 0; j < len(arr); j++ {
            digit := arr[j]
            digit *= i
            digit += rem
            arr[j] = digit%10
            rem = digit/10
        }
        for rem != 0 {
            arr = append(arr, rem%10)
            rem /= 10
        }
    }
    for i := len(arr)-1; i > -1; i-- {
        fmt.Print(arr[i])
    }
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int(nTemp)

    extraLongFactorials(n)
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
