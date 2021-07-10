package utils

import "strings"

//ReverseString 反转字符
func ReverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

//GetDIValue [00001111 01110000] 去中括号，重新排序 11110000+00001110
func GetDIValue(input string) string {
	value := strings.Split(string(input[1:len(input)-1]), " ")
	return ReverseString(value[0]) + ReverseString(value[1])
}
