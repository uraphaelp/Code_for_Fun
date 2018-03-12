//将字符串中的 " " 替换成 "%20"

import "strings"

//do not apply extra space
func Replacespace1(s string) string {
	s=strings.Replace(s, " ", "%20", -1) /* -1 代表替换 s 中的所有目标字符串*/
	return s
}

// no use of any package
func Replacespace2(str string) {
/*this method has to apply new space
pay attention to annotation, that's standard method
 */
	/*determine how many spaces
	count:=0
	for i:=0; i<len(str); i++ {
		if str[i]==' ' {
			count++
		}
	}
	*/
	var newstr string
	for i:=len(str)-1; i>=0; i-- {
		if str[i]!=' ' {
			newstr=string(str[i])+newstr
			/*
			count--;
			str[i+2*count]='%';
			str[i+2*count+1]='2';
			str[i+2*count+2]='0';
			*/
		} else{
			newstr="%20"+newstr
		}
	}
	return newstr
}
