package main

import (
	"bytes"
	"container/list"
	"strconv"
	"strings"
)

type node struct {
	count  int
	value  string
	isLeaf bool
	next   []*node
	pre    *node
}

func decodeStringTree(s string) string {
	if len(s) < 1 {
		return ""
	}
	num := ""
	var head *node
	var next *node
	root := &node{count: 1}
	head = root
	contents := ""
	for _, v := range s {
		switch {
		case v >= '0' && v <= '9':
			num = strings.Join([]string{num, string(v)}, "")
		case v == '[':
			if len(contents) > 0 {
				head.next = append(head.next, &node{count: 1, value: contents, isLeaf: true})
				contents = ""
			}
			c, _ := strconv.Atoi(num)
			num = ""
			next = &node{count: c}
			head.next = append(head.next, next)
			next.pre = head
			head = next
		case v == ']':
			if len(contents) > 0 {
				head.next = append(head.next, &node{count: 1, value: contents, isLeaf: true})
				contents = ""
			}
			head = head.pre
		default:
			contents = strings.Join([]string{contents, string(v)}, "")
		}
	}
	if s[len(s)-1] != ']' {
		head.next = append(head.next, &node{count: 1, value: contents, isLeaf: true})
	}

	var array bytes.Buffer
	var nodeHelper func(*node) []byte
	nodeHelper = func(head *node) []byte {

		var buffer bytes.Buffer
		if head.isLeaf {
			buffer.WriteString(head.value)
		} else {
			for i := 0; i < head.count; i++ {
				for _, v := range head.next {
					sub := nodeHelper(v)
					buffer.Write(sub)
				}
			}
		}
		return buffer.Bytes()
	}
	array.Write(nodeHelper(root))
	return array.String()
}

// 辅助栈 和 递归法
func decodeString(s string) string {
	multi := 0
	contents := ""
	stack := list.New()
	type res struct {
		m   int
		str string
	}
	for _, v := range s {
		switch {
		case v >= '0' && v <= '9':
			multi = multi*10 + int(v-'0')
		case v == '[':
			stack.PushBack(res{m: multi, str: contents})
			multi = 0
			contents = ""
		case v == ']':
			last := stack.Remove(stack.Back()).(res)
			buff := bytes.Buffer{}
			buff.WriteString(last.str)
			for i := 0; i < last.m; i++ {
				buff.WriteString(contents)
			}
			contents = buff.String()
		default:
			contents = strings.Join([]string{contents, string(v)}, "")
		}
	}
	return contents
}

func main() {
	println(decodeString("3[a]2[bc]"))
}
