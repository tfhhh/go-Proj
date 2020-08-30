package main

import (
	"fmt"
	"strings"
)
/*
最小公共前缀：
输入: ["flower","flow","flight"]
输出: "fl"
 */
func longestCommonPrefix(str []string) string{
	if len(str) < 1 {
		return ""
	}
	refer := str[0]	//flower
	for _, k := range str {
		for strings.Index(k,refer) != 0{	//strings.Index(str,"s"),返回s在str中的位置,不存在为-1
			if len(refer) == 0{
				return ""
			}else {
				refer = refer[:len(refer)-1]	//从最后一个字母依次往前截取
			}
		}
	}
	return refer
}
func main() {
	a :=[] string{"flower","flow","flight"}
	fmt.Println(longestCommonPrefix(a))

	//化繁为简
	b := "string"
	c := "stredf"
	for strings.Index(c,b) != 0 { //strings.Index(str,"s"),返回s在str中的位置,不存在为-1
		if len(b) == 0 {
			fmt.Println("")
		} else {
			b = b[:len(b)-1] //从最后一个字母依次往前截取
		}
		fmt.Println(b) //依次输出strin，stri，str
		fmt.Println("index:",strings.Index(c,b)) //-1,-1,0
	}
	//理解strings.Index(str,str)->int
	d := "string"
	e := "stredf"
	fmt.Println(strings.Index(e,d[:4]))	//-1
	fmt.Println(strings.Index(e,d[:3]))	//0
	fmt.Println(strings.Index(e,d[:2]))	//0
}
