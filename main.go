package main

import "fmt"

const (
	maxInt32 = 1<<31 - 1
	minInt32 = -1 << 31
)

// Atoi parses an integer the way C's atoi contract is usually specified in
// interviews: skip leading spaces, accept one optional sign, consume digits
// until a non-digit, and clamp to the 32-bit signed range.
func Atoi(s string) int {
	i, n := 0, len(s)
	for i < n && s[i] == ' ' {
		i++
	}
	sign := 1
	if i < n && (s[i] == '+' || s[i] == '-') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}
	total := 0
	for i < n && s[i] >= '0' && s[i] <= '9' {
		total = total*10 + int(s[i]-'0')
		if sign == 1 && total > maxInt32 {
			return maxInt32
		}
		if sign == -1 && -total < minInt32 {
			return minInt32
		}
		i++
	}
	return sign * total
}

// StrStr returns the index of the first occurrence of needle in haystack,
// -1 if absent, 0 for an empty needle. Simple quadratic scan — the linear
// upgrades (KMP, Z) live in their own repos later in the sheet.
func StrStr(haystack, needle string) int {
	n, m := len(haystack), len(needle)
	if m == 0 {
		return 0
	}
	for i := 0; i+m <= n; i++ {
		if haystack[i:i+m] == needle {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(Atoi("   -42abc"), StrStr("sadbutsad", "sad"))
}
