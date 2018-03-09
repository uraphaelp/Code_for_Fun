## 在二维数组中寻找target

## 算法分析
**从 左下角/右上角/左下角开始找：**

这几个位置共同特点(以左下角为例)：
```
for condition:
    if target>nums[i][j] {
        j++
    } else if target<nums[i][j] {
        i--
    } else {
        return true
    }
}
```
