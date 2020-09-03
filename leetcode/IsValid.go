package main

import "fmt"
//字符串里的括号是否匹配
func isValid(s string) bool {
	if s == "" {
		return true
	}
	res := make([]string,len(s))
	var top = 0

	for i := 0; i < len(s); i++ {
		//进栈,栈为空或者字符串元素与栈顶不匹配则进栈
		if res[0] == "" {
			res[top] = string(s[i])
			top++
			continue	//每个元素只执行一次
		}
		if res[top-1] == "(" && string(s[i]) != ")" {
			res[top] = string(s[i])
			top++
			continue
		}
		if res[top-1] == "[" && string(s[i]) != "]" {
			res[top] = string(s[i])
			top++
			continue
		}
		if res[top-1] == "{" && string(s[i]) != "}" {
			res[top] = string(s[i])
			top++
			continue
		}
		//出栈，栈顶与元素匹配则出栈
		if res[top-1] == "(" && string(s[i]) == ")" {
			top--
			res[top] = ""
			continue
		}
		if res[top-1] == "[" && string(s[i]) == "]" {
			top--
			res[top] = ""
			continue
		}
		if res[top-1] == "{" && string(s[i]) == "}" {
			top--
			res[top] = ""
			continue
		}
	}
	return top==0
}
func main() {
	a:= "()[]{}"
	b := "){"
	c := "()"
	d := "(}{)"
	e := "((({[]})))"
	fmt.Println(isValid(a))
	fmt.Println(isValid(b))
	fmt.Println(isValid(c))
	fmt.Println(isValid(d))
	fmt.Println(isValid(e))
	/*
	true
	false
	true
	false
	true

	Process finished with exit code 0
	 */

}
