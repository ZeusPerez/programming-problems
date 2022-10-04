package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'treeConstruction' function below.
 *
 * The function is expected to return a 2D_INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER N
 *  2. LONG_INTEGER X
 */

func treeConstruction(N int32, X int64) [][]int32 {
	if int32(X) < N-1 || X < 1 || N <= 1 {
		result := [][]int32{{-1, -1}}
		return result
	}

	var result [][]int32
	distance := int64(1)
	parents := []int32{0, 1}
	parent := int32(1)
	for i := 2; int32(i) <= N; i++ {
		pendingNodes := N - int32(i)
		for int32(X-distance) < pendingNodes {
			parent = parents[parent]
			distance--
		}
		X = X - distance
		result = append(result, []int32{parent, int32(i)})
		parents = append(parents, parent)
		parent = int32(i)
		distance++

	}
	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	// fileReader, err := os.Open("./in.txt")
	// checkError(err)

	// reader := bufio.NewReaderSize(fileReader, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	TTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	T := int32(TTemp)

	for TItr := 0; TItr < int(T); TItr++ {
		firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		NTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
		checkError(err)
		N := int32(NTemp)

		X, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
		checkError(err)

		result := treeConstruction(N, X)

		for i, rowItem := range result {
			for j, colItem := range rowItem {
				fmt.Fprintf(writer, "%d", colItem)

				if j != len(rowItem)-1 {
					fmt.Fprintf(writer, " ")
				}
			}

			if i != len(result)-1 {
				fmt.Fprintf(writer, "\n")
			}
		}

		fmt.Fprintf(writer, "\n")
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
