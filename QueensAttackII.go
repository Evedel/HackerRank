package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

func isObs(obs map[int]map[int]int, i,j int) bool {
    if _, ok1 := obs[i]; ok1 {
        if _, ok2 := obs[i][j]; ok2 {
            return true
        }
    }
    return false
}
// Complete the queensAttack function below.
// 
// You will be given a square chess board with one queen
//    and a number of obstacles placed on it.
// Determine how many squares the queen can attack.
// 
// A queen is standing on an NxN chessboard.
// The chess board's rows are numbered from 1 to n, going from bottom to top.
// Its columns are numbered from 1 to n, going from left to right.
// Each square is referenced by a tuple (r,c),
// describing the row, r, and column, c, where the square is located.
// 
// The queen is standing at position (r_q,c_q).
// In a single move, she can attack any square in any of the eight directions
// (left, right, up, down, and the four diagonals).
// 
// There are obstacles on the chessboard,
// each preventing the queen from attacking any square beyond it on that path.
// 
// Given the queen's position and the locations of all the obstacles,
// find and print the number of squares the queen can attack from her position.
// 
// O(n)
// 
// https://www.hackerrank.com/challenges/queens-attack-2
// 
func queensAttack(n int, k int, r_q int, c_q int, obstacles [][]int) int {
    obs := map[int]map[int]int{}
    for i := range(obstacles) {
        if _, ok := obs[obstacles[i][0]-1]; ok {
            obs[obstacles[i][0]-1][obstacles[i][1]-1] = -1
        } else {
            obs[obstacles[i][0]-1] = map[int]int{}
            obs[obstacles[i][0]-1][obstacles[i][1]-1] = -1
        }
    }
    rq := r_q-1
    rc := c_q-1

    turns := 0
    // all in the column
    for i := rq+1; i < n; i++ {
        if !isObs(obs,i,rc) {
            turns += 1
        } else {
            break
        }
    }
    for i := rq-1; i >= 0; i-- {
        if !isObs(obs,i,rc) {
            turns += 1
        } else {
            break
        }
    }
    // row
    for j := rc+1; j < n; j++ {
        if !isObs(obs,rq,j) {
            turns += 1
        } else {
            break
        }
    }
    for j := rc-1; j >= 0; j-- {
        if !isObs(obs,rq,j) {
            turns += 1
        } else {
            break
        }
    }
    // diagonal 1
    k = 1
    for ((rq+k) < n) && ((rc+k) < n) {
        if !isObs(obs,rq+k,rc+k) {
            turns += 1
        } else {
            break
        }
        k++
    }
    k = 1
    for ((rq-k) > -1) && ((rc-k) > -1) {
        if !isObs(obs,rq-k,rc-k) {
            turns += 1
        } else {
            break
        }
        k++
    }
    // diagonal 2
    k = 1
    for ((rq+k) < n) && ((rc-k) > -1) {
        if !isObs(obs,rq+k,rc-k) {
            turns += 1
        } else {
            break
        }
        k++
    }
    k = 1
    for ((rq-k) > -1) && ((rc+k) < n) {
        if !isObs(obs,rq-k,rc+k) {
            turns += 1
        } else {
            break
        }
        k++
    }

    // for i := 0; i < n; i++ {
    //     for j := 0; j < n; j++ {
    //         fmt.Printf("%4d", grid[i][j])
    //     }
    //     fmt.Println()
    // }

    return turns
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nk := strings.Split(readLine(reader), " ")

    nTemp, err := strconv.ParseInt(nk[0], 10, 64)
    checkError(err)
    n := int(nTemp)

    kTemp, err := strconv.ParseInt(nk[1], 10, 64)
    checkError(err)
    k := int(kTemp)

    r_qC_q := strings.Split(readLine(reader), " ")

    r_qTemp, err := strconv.ParseInt(r_qC_q[0], 10, 64)
    checkError(err)
    r_q := int(r_qTemp)

    c_qTemp, err := strconv.ParseInt(r_qC_q[1], 10, 64)
    checkError(err)
    c_q := int(c_qTemp)

    var obstacles [][]int
    for i := 0; i < int(k); i++ {
        obstaclesRowTemp := strings.Split(readLine(reader), " ")

        var obstaclesRow []int
        for _, obstaclesRowItem := range obstaclesRowTemp {
            obstaclesItemTemp, err := strconv.ParseInt(obstaclesRowItem, 10, 64)
            checkError(err)
            obstaclesItem := int(obstaclesItemTemp)
            obstaclesRow = append(obstaclesRow, obstaclesItem)
        }

        if len(obstaclesRow) != int(2) {
            panic("Bad input")
        }

        obstacles = append(obstacles, obstaclesRow)
    }

    result := queensAttack(n, k, r_q, c_q, obstacles)

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
