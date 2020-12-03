package util

/**
分页工具
*/
func PageUtil(count int, pageSize int, pageNum int) map[string]interface{} {
	pageMap := make(map[string]interface{})

	//总条数
	pageMap["count"] = count
	//	每页条数
	if pageSize <= 0 {
		pageSize = 1
	}
	pageMap["pageSize"] = pageSize
	//当前页数
	if pageNum < 0 {
		pageNum = 0
	}
	pageMap["pageNum"] = pageNum
	//总页数
	pageMap["pages"] = func(count int, pageSize int) int {
		if count <= 0 {
			return 0
		} else if count < pageSize {
			return 1
		} else {
			if count%pageSize > 0 {
				return (count / pageSize) + 1
			}
			return count / pageSize
		}
	}(count, pageSize)
	//页码标签
	pageMap["pageData"] = func() []int {
		p := pageMap["pageNum"]
		p2 := pageMap["pages"]
		//当前页码标签前后各显示多少页码标签
		size := 2
		d := make([]int, 0)
		if p.(int) <= 0 {
			return d
		} else {
			//取前几位
			for a := size; a >= 1; a-- {
				b := p.(int) - a
				if b >= 1 {
					d = append(d, b)
				} else {
					continue
				}
			}
			//当前页
			d = append(d, p.(int))
			//	取后几位
			for a := 1; a <= size; a++ {
				b := p.(int) + a
				if b <= p2.(int) {
					d = append(d, b)
				} else {
					continue
				}
			}
		}
		return d
	}()

	return pageMap
}
