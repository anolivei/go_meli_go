package ordering

import "sort"

func SortNums(num []int) []int {
	sort.Ints(num)
	return num
}
