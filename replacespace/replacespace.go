//将字符串中的 " " 替换成 "%20"

import "strings"

//do not apply extra space
func Replacespace1(s string) string {
	s=strings.Replace(s, " ", "%20", -1)
	return s
}

// no use of any package
func Replacespace2(s string) string {
}
