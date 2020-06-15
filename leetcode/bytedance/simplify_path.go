package main

import "fmt"

func simplifyPath(path string) string {
	dir := make([]string, 0)
	word := ""
	path2 := fmt.Sprintf("%s/", path)
	for _, w := range path2 {
		if w == '/' {
			if word != "" {
				if word == ".." {
					if len(dir) > 0 {
						dir = dir[:len(dir)-1]
					}
				} else if word != "." {
					dir = append(dir, word)
				}
			}
			word = ""
		} else {
			word = fmt.Sprintf("%s%s", word, string(w))
		}
	}

	if len(dir) > 0 {
		result := ""
		for _, w := range dir {
			result = fmt.Sprintf("%s/%s", result, w)
		}
		return result
	}
	return "/"
}

func main() {
	print(simplifyPath("/a//b////c/d//././/.."))
}
