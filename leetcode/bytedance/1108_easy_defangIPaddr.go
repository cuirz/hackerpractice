package main

import "strings"

//1108. IP 地址无效化
//给你一个有效的 IPv4 地址address，返回这个 IP 地址的无效化版本。
//
//所谓无效化IP 地址，其实就是用"[.]"代替了每个 "."。
//
//
//
//示例 1：
//
//输入：address = "1.1.1.1"
//输出："1[.]1[.]1[.]1"
//示例 2：
//
//输入：address = "255.100.50.0"
//输出："255[.]100[.]50[.]0"
//
//
//提示：
//
//给出的address是一个有效的 IPv4 地址

func defangIPaddr(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}
