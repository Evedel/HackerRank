package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
    "math"
)

// Complete the encryption function below.
// An English text needs to be encrypted using the following encryption scheme.
// First, the spaces are removed from the text.
// Let "L" be the length of this text.
// Then, characters are written into a grid,
// whose rows and columns have the following constraints:
//      floor(sqrt(L)) <= row <= column <= ceil(sqrt(L))
// Ensure that row*column > L
// If multiple grids satisfy the above conditions,
// choose the one with the minimum area, i.e. min(row*column).
// 
// For example, the sentence
// s="if man was meant to stay on the ground god would have given us roots",
// after removing spaces is 54 characters long.
// sqrt(54) is between 7 and 8,
// so it is written in the form of a grid with 7 rows and 8 columns.
// 
// ifmanwas  
// meanttos          
// tayonthe  
// groundgo  
// dwouldha  
// vegivenu  
// sroots
// 
// The encoded message is obtained by displaying the characters in a column,
// inserting a space, and then displaying the next column and inserting a space,
// and so on. For example, the encoded message for the above rectangle is:
// 
// imtgdvs fearwer mayoogo anouuio ntnnlvt wttddes aohghn sseoau
// 
// You will be given a message to encode and print.
// 
// https://www.hackerrank.com/challenges/encryption
// 
// O(n)
// 
func encryption(s string) string {
    s = strings.Join(strings.Split(s," "),"")
    L := len(s)
    sL := math.Sqrt(float64(L))
    s1 := int(sL)
    s2 := s1
    if (float64(s1) != sL) { s2 = s1+1 }
    if (s1*s2 < L) { s1 = s2 }
    encs := ""
    orgi := 0
    shft := 0
    for i := 0; i < L; i++ {
        encs += string(s[orgi])
        orgi += s2
        if (orgi >= L) {
            encs += " "
            shft += 1
            orgi = shft
        }
    }
    return encs
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    s := readLine(reader)

    result := encryption(s)

    fmt.Fprintf(writer, "%s\n", result)

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
