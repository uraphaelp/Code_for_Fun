package hexadecimal

import (
	"testing"
	"fmt"
)

func TestHexadecimal36(t *testing.T) {
	testcase:=[]string{"0", "2", "9", "18", "z", "1y"}
	for i:=0; i<len(testcase); i++ {
		for j:=i+1; j<len(testcase); j++ {
			fmt.Println(Hexadecimal36(testcase[i], testcase[j]))
		}
	}
}

//Output:[2, 9, 18, z, 1y, b, 1a, 11, 20, 1h, 18, 27, 27, 36, 2x]
