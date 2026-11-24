package main

import "testing"

func TestAtoi(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"42", 42},
		{"   -42", -42},
		{"4193 with words", 4193},
		{"words and 987", 0},
		{"-91283472332", -2147483648}, // clamped
		{"91283472332", 2147483647},   // clamped
		{"+1", 1},
		{"", 0},
		{"+-12", 0}, // only one sign allowed
		{"2147483648", 2147483647},
		{"-2147483648", -2147483648},
	}
	for _, c := range cases {
		if got := Atoi(c.in); got != c.want {
			t.Errorf("Atoi(%q) = %d, want %d", c.in, got, c.want)
		}
	}
}

func TestStrStr(t *testing.T) {
	cases := []struct {
		h, n string
		want int
	}{
		{"sadbutsad", "sad", 0},
		{"leetcode", "leeto", -1},
		{"hello", "ll", 2},
		{"abc", "", 0},
		{"", "a", -1},
		{"aaa", "aaaa", -1},
		{"mississippi", "issip", 4},
	}
	for _, c := range cases {
		if got := StrStr(c.h, c.n); got != c.want {
			t.Errorf("StrStr(%q,%q) = %d, want %d", c.h, c.n, got, c.want)
		}
	}
}
