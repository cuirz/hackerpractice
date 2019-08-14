package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Complete the appendAndDelete function below.
func appendAndDelete(s string, t string, k int32) string {
	ls := len(s)
	lt := len(t)
	size := 0
	index := 0
	if ls >= lt {
		size = lt
	} else {
		size = ls
	}
	for i := 0; i < size; i++ {
		if s[i] == t[i] {
			index++
		} else {
			break
		}
	}
	lst := ls + lt
	tmp := int(k) - (lst - 2*index)
	fmt.Println(tmp, lst, index)
	if int(k) > lst || tmp == 0 || (tmp > 0 && tmp%2 == 0) {
		return "Yes"
	}
	return "No"
	//qwerasdf
	//qwerbsdf
	//6

	//aaaaaaaaaa
	//aaaaa
	//7

	//abcd
	//abcdert
	//10

}

func main() {
	//reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	//s := readLine(reader)

	//t := readLine(reader)

	//kTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	//checkError(err)
	//k := int32(kTemp)

	//result := appendAndDelete(s, t, k)
	//result := appendAndDelete("aaaaaaaaaa", "aaaaa", 7)
	//result := appendAndDelete("qwerasdf", "qwerbsdf", 6)
	result := appendAndDelete("aba", "aba", 7)
	//result := appendAndDelete("abcd", "abcdert", 10)

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
