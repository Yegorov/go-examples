package tasks

// https://solvit.space/coding/3347
// https://algo.monster/liteproblems/394
// Декодирование строки

import (
	"strconv"
	"strings"
)

func DecodeStringWithoutNesting(s string) string {
	res := make([]rune, 0, 300)
	sub := make([]rune, 0, 300)
	nums := make([]rune, 0, 10)
	digits := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	n := 0
	ss := []rune(s)
	for _, c := range ss {
		switch {
		case Contains(digits, c):
			nums = append(nums, c)
		case c == '[':
			n, _ = strconv.Atoi(string(nums))
		case c == ']':
			for i := 0; i < n; i++ {
				for j := 0; j < len(sub); j++ {
					res = append(res, sub[j])
				}
			}
			n = 0
			nums = nums[0:0]
			sub = sub[0:0]
		default:
			if n == 0 {
				res = append(res, c)
			} else {
				sub = append(sub, c)
			}
		}
	}
	return string(res)
}

func Contains(s []rune, c rune) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			return true
		}
	}
	// or use for with range
	// for _, cc := range s {
	// 	if cc == c {
	// 		return true
	// 	}
	// }
	//
	// or use slices.Contains
	return false
}

func DecodeString(s string) string {
	st1 := make([]int, 0, 100)
	st2 := make([]string, 0, 100)
	str := ""
	k := 0
	for _, c := range s {
		if Contains([]rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}, c) {
			k = k*10 + int(c) - int('0')
		} else if c == '[' {
			st1 = append(st1, k)
			st2 = append(st2, str)
			str = ""
			k = 0
		} else if c == ']' {
			nn := st1[len(st1)-1]
			st1 = st1[:len(st1)-1]
			ss := st2[len(st2)-1]
			st2 = st2[:len(st2)-1]
			rs := strings.Builder{}
			for range nn {
				rs.WriteString(str)
			}
			str = ss + rs.String()
		} else {
			str += string(c)
		}
	}
	return str
}
