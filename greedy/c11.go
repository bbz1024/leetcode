package greedy

import (
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	// 倒序排序
	sort.Slice(people, func(i, j int) bool {
		//	先比较身高，再是个数
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})
	// 前往后调整
	for i := 0; i < len(people); i++ {
		// 找到插入位置
		// 移动到[i][1]的位置
		insertIdx := people[i][1]
		cur := i
		temp := people[i]
		for cur > insertIdx {
			people[cur] = people[cur-1]
			cur--
		}
		people[cur] = temp
		// 写回
	}
	return people
}

/*

[[7 0] ,[7 1], [6 1] [5 0] [5 2] [4 4]]

[[7 0] , [6 1],[7 1] [5 0] [5 2] [4 4]]

[[7 0] , [6 1],[7 1] [5 0] [5 2] [4 4]]

*/
