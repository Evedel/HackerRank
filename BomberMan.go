package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

func printGrid(gr [][]int) {
    for i := range(gr) {
        for j := range gr[i] {
            // fmt.Printf("%4d",gr[i][j])
            if gr[i][j] == -1 {
                fmt.Print(".")
            } else {
                fmt.Print("O")
            }
        }
        fmt.Print("\n")
    }
}

// Complete the bomberMan function below.
// Bomberman lives in a rectangular grid.
// Each cell in the grid either contains a bomb or nothing at all.

// Each bomb can be planted in any cell of the grid but once planted,
// it will detonate after exactly 3 seconds. Once a bomb detonates,
// it's destroyed — along with anything in its four neighboring cells.
// This means that if a bomb detonates in cell [i,j],
// any valid cells [i+/-1,j] and [i,j+/-1] are cleared.
// If there is a bomb in a neighboring cell,
// the neighboring bomb is destroyed without detonating,
// so there's no chain reaction.

// Bomberman is immune to bombs, so he can move freely throughout the grid.
// Here's what he does:

// 1) Initially, Bomberman arbitrarily plants bombs in some of the cells,
//    the initial state.
// 2) After one second, Bomberman does nothing.
// 3) After one more second, Bomberman plants bombs in all cells without bombs,
//    thus filling the whole grid with bombs. No bombs detonate at this point.
// 4) After one more second, any bombs planted exactly three seconds ago will detonate.
//    Here, Bomberman stands back and observes.
// Bomberman then repeats steps 3 and 4 indefinitely.
// Note that during every second Bomberman plants bombs,
// the bombs are planted simultaneously (i.e., at the exact same moment),
// and any bombs planted at the same time will detonate at the same time.

// Given the initial configuration of the grid with the locations of Bomberman's
// first batch of planted bombs, determine the state of the grid after
// 1 < n < 10^9 seconds.
// 
// https://www.hackerrank.com/challenges/bomber-man/problem
// 
func bomberMan(n int, grid []string) []string {
    gr := [][]int{}
    for i := range(grid) {
        gr = append(gr, []int{})
        for _, c := range grid[i] {
            if c == '.' {
                gr[i] = append(gr[i], -1) 
            } else {
                gr[i] = append(gr[i],  0)
            }
        }
    }
    // initial state --- state after 0 seconds
    cn := 0
    // bomberman does nothing --- state after 1 seconds
    cn = 1
    if n > 5 {
        n = n%4 + 4
    }
    // there exist only 4 _stable_ states,
    // initial -> planted #1 all field -> initial exploded
    // -> planted #2 all field -> planted #1 exploded
    // -> planted #1 all field => repeat.
    // Which is why we can simmulate only first stable state
    for cn < n {
        cn += 1
        if (cn%2 == 0) {
            for i := range(gr) {
                for j := range gr[i] {
                    if gr[i][j] == -1 { gr[i][j] = cn}
                }
            }
        } else {
            li := len(gr)-1
            lj := len(gr[0])-1
            for i := range(gr) {
                for j := range gr[i] {
                    if (cn - gr[i][j]) == 3 {
                        if (i > 0)  { if gr[i-1][j] != gr[i][j] { gr[i-1][j] = -1 } }
                        if (j > 0)  { if gr[i][j-1] != gr[i][j] { gr[i][j-1] = -1 } }
                        if (i < li) { if gr[i+1][j] != gr[i][j] { gr[i+1][j] = -1 } }
                        if (j < lj) { if gr[i][j+1] != gr[i][j] { gr[i][j+1] = -1 } }
                        gr[i][j] = -1
                    }
                }
            }
        }
    }
    newGrid := []string{}
    for i := range(gr) {
        newStr := ""
        for j := range gr[i] {
            if (gr[i][j] == -1) {
                newStr += "."
            } else {
                newStr += "O"
            }
        }
        newGrid = append(newGrid, newStr)
    }

    return newGrid
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    rcn := strings.Split(readLine(reader), " ")

    rTemp, err := strconv.ParseInt(rcn[0], 10, 64)
    checkError(err)
    r := int32(rTemp)

    // cTemp, err := strconv.ParseInt(rcn[1], 10, 64)
    // checkError(err)
    // c := int32(cTemp)

    nTemp, err := strconv.ParseInt(rcn[2], 10, 64)
    checkError(err)
    n := int(nTemp)

    var grid []string

    for i := 0; i < int(r); i++ {
        gridItem := readLine(reader)
        grid = append(grid, gridItem)
    }

    result := bomberMan(n, grid)

    for i, resultItem := range result {
        fmt.Fprintf(writer, "%s", resultItem)

        if i != len(result) - 1 {
            fmt.Fprintf(writer, "\n")
        }
    }

    fmt.Fprintf(writer, "\n")

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
