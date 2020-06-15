package main

import (
	"fmt"
	"strconv"
	"strings"
)

//输入: "25525511135"
//输出: ["255.255.11.135", "255.255.111.35"]

func restoreIpAddresses(s string) []string {
	size := len(s)
	if size < 4 || size > 12 {
		return []string{}
	}
	result := new([]string)
	var array []string
	splitIP(s, result, array, 0)
	return *result
}

func splitIP(s string, result *[]string, array []string, count int) []string {
	if len(s) == 0 && count < 4 {
		return array[:len(array)-1]
	} else if count > 3 {
		if len(s) == 0 {
			*result = append(*result, strings.Join(array, "."))
		}
		return array[:len(array)-1]

	}
	for i := 3; i > 0; i-- {
		if i > len(s) {
			continue
		}
		num := s[:i]
		if validIp(num) {
			array = append(array, num)
			array = splitIP(s[i:], result, array, count+1)
		}

	}
	if len(array) < 1 {
		return []string{}
	}
	return array[:len(array)-1]
}

func validIp(ip string) bool {
	if len(ip) > 1 && ip[0] == '0' {
		return false
	}
	_ip, _ := strconv.ParseInt(ip, 10, 64)
	if 0 <= _ip && _ip < 256 {
		return true
	}
	return false
}

func main() {
	println(fmt.Sprintf("%v", restoreIpAddresses("25525511135")))
}
