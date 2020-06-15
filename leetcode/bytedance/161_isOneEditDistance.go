package main

func isOneEditDistance(s string, t string) bool {
	n, m := len(s), len(t)
	if n == 0 && m == 0 {
		return false
	}
	if n < m {
		return isOneEditDistance(t, s)
	}
	if n == 1 && m == 0 {
		return true
	}
	if n-m > 1 {
		return false
	}

	for i := 0; i < m; i++ {
		if s[i] != t[i] {
			if n == m {
				return s[i+1:] == t[i+1:]
			}
			return s[i+1:] == t[i:]
		}
	}

	return n > m
}
