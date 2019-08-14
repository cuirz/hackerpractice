package main

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
	"os"
	"strconv"
	"strings"
)

// Complete the extraLongFactorials function below.
func extraLongFactorials(n int32) {
	large := big.NewInt(1)
	for i := int32(2); i <= n; i++ {
		large = large.Mul(large, big.NewInt(int64(i)))
	}
	fmt.Println(large.String())
}

func extraLongFactorials2(n int32) {
	large := big.NewInt(1)
	for i := int32(2); i <= n; i++ {
		large = large.Mul(large, big.NewInt(int64(i)))
	}
	fmt.Println(large.String())
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

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
