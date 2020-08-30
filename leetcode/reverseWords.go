package main

import (
	"fmt"
	"strings"
)

/*
输入："Let's take LeetCode contest"
输出："s'teL ekat edoCteeL tsetnoc"
 */
func reverseWords(s string) string {
	s = " "+s+" "
	var a int
	var str string
	for i := 0; i < len(s); i++ {
		var temp string
		if s[i] == 32 || i == len(s)-1{
			temp = s[a:i+1]
			for j := len(temp)-1; j >= 0 ; j-- {
				str = str + string(temp[j])
			}
			a = i
		}
	}
	str = strings.Replace(str, "  ", " ", -1)
	return str[1:len(str)-1]
}
func main() {
	s := "Let's take LeetCode contest"
	str := reverseWords(s)
	fmt.Println(s)
	fmt.Println(str)
}