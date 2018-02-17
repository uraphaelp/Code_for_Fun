//36进制加法：0-9，a-z
package hexadecimal

import "strconv"

//use non-capital
func Hexadecimal36(num1 string, num2 string) string {
	var (
		result string
		carry int
		i=len(num1)-1
		j=len(num2)-1
	)
	for i>=0 || j>=0 || carry!=0 {
		var x, y int
    //先把每一位数转换了
		if i>=0 {
			x=Convert(num1[i])
			i--
		} else {
			x=0
		}
		if j>=0 {
			y=Convert(num2[j])
			j--
		} else {
			y=0
		}
    
    //不需要进位的数+数的情况
		if x+y+carry<10 {
			result+=strconv.Itoa(x+y+carry)
			carry=0
    //不需要进位的数+字母的情况
		} else if x+y+carry<=35 && x+y+carry>=10 {
			result+=string('a'+(x+y+carry-10))
			carry=0
    //本位进位后为数字的情况
		} else if x+y+carry>35 && (x+y+carry)%36<10 {
			result+=strconv.Itoa((x+y+carry)%36)
    //注意进位的表示方法
			carry=(x+y+carry)/36
    //本位进位后为字母的情况
		} else if x+y+carry>35 && (x+y+carry)%36>=10 {
			result+=string('a'+((x+y+carry)%36)-10)
			carry=(x+y+carry)/36
		}
	}
	return Reverse(result)
}

//单个位数转换为十进制数
func Convert(origin byte) int {
	var goal int
	if origin>=49 && origin<=57 {
		goal=int(origin-'0')
	} else if origin>=97 && origin<=122 {
		goal=int(origin-'a'+10)
	} else {
		goal=0
	}
	return goal
}

//逆转数组
func Reverse(s string) string {
	var result string
	for i:=len(s)-1; i>=0; i-- {
		result+=string(s[i])
	}
	return result
}
