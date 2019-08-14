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
 * Complete the 'nonDivisibleSubset' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER k
 *  2. INTEGER_ARRAY s
 */
//counts = [0] * k
//for number in numbers:
//counts[number % k] += 1
//
//count = min(counts[0], 1)
//for i in range(1, k//2+1):
//if i != k - i:
//count += max(counts[i], counts[k-i])
//if k % 2 == 0:
//count += 1
//
//print count
func nonDivisibleSubset(k int32, s []int32) int32 {
	// Write your code here
	min := func(a, b int32) int32 {
		if a > b {
			return b
		}
		return a
	}
	max := func(a, b int32) int32 {
		if a > b {
			return a
		}
		return b
	}
	counts := make(map[int32]int32)
	for _, v := range s {
		counts[v%k] += 1
	}
	count := min(counts[0], 1)

	size := int32(k / 2)
	for i := int32(1); i <= size; i++ {
		if i != k-i {
			count += max(counts[i], counts[k-i])
		}

	}
	if k%2 == 0 {
		count += 1
	}
	println(counts)
	fmt.Println(counts)
	return count

	//subset := make(map[int32]int32)
	//for i, v := range s {
	//	for _, ov := range s[i+1:] {
	//		if (v + ov )% k != 0 {
	//			subset[v] += 1
	//			subset[ov] += 1
	//			//println(v,ov,v+ov)
	//		}
	//	}
	//
	//}
	//for _,v := range subset{
	//	print(v," ")
	//}
	//println("")
	//
	//return int32(len(subset))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	sTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var s []int32

	for i := 0; i < int(n); i++ {
		sItemTemp, err := strconv.ParseInt(sTemp[i], 10, 64)
		checkError(err)
		sItem := int32(sItemTemp)
		s = append(s, sItem)
	}

	result := nonDivisibleSubset(k, s)

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
