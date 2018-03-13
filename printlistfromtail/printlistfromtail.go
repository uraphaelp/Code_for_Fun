//从尾到头打印链表
type ListNode struct {
	Val int
	Next *ListNode
}

type stack struct {
	Top *ListNode
	Bottom *ListNode
}

//in recursive way
func Printlistfromtailinrecur(head *ListNode) []int {
	var res []int
	p:=head
	if p!=nil {
		Printlistfromtailinrecur(p)
		res=append(res, p.Val)
	}
	return res
}

//use stack
func Printlistfromtailinstack(head *ListNode) []int {
	var res []int
	p:=head
	stack:=new(stack)
	for p!=nil {
		stack.Push(p)
		p=p.Next
	}
	for !stack.Isempty() {
		res=append(res, stack.Pop().Val)
	}
	return res
}
