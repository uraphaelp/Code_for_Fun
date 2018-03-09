func Findele(nums [][]int, target int) (bool, int) {
	var count int
	l:=len(nums)-1
	j:=0
	for l>=0 && j<=len(nums[0])-1 {
		if nums[l][j]>target {
			l--
			count++
		} else if nums[l][j]<target {
			j++
			count++
		} else {
			return true, count
		}
	}
	return false, count
}

//count:计数寻找次数
