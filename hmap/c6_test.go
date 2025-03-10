package hmap

import (
	"fmt"
	"testing"
)

func TestFourSum(t *testing.T) {
	//	[1,-2,-5,-4,-3,3,3,5]
	nums := []int{1, -2, -5, -4, -3, 3, 3, 5}
	target := -11
	res := fourSum(nums, target)
	fmt.Println(res)
}

/*

-2,-1,0,0,1,2
*/
