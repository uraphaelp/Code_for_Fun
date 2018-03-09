import (
	"testing"
	"fmt"
)

func TestFindele(t *testing.T) {
	testcase:=[][]int{{1, 3, 8}, {2, 7, 10}, {6, 9, 15}}
	fmt.Println(Findele(testcase, 8))
	fmt.Println(Findele(testcase, 4))
}

/*OUTPUT:
true, 4
false, 5
*/
