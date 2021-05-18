package series

// All returns all substrings
func All(n int, s string) []string {
	var res []string
	for i := 0; i <= len(s) - n; i++ {
		res = append(res, s[i:i+n])
	}
	return res
}

// UnsafeFirst returns the first substring
func UnsafeFirst(n int, s string) string {
	if len(s) < n {
		return ""
	}

	return s[0:n]
}

// First returns the first substring
func First(n int, s string) (string, bool) {
	if len(s) < n {
		return "", false
	}

	return s[0:n], true
}
