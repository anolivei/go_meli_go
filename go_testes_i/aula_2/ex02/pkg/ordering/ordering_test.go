package ordering

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSortNums(t *testing.T) {
	num := []int{1, 6, 8, 5, 3, 2, 9, 4, 7}
	result := SortNums(num)
	expectedResult := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	assert.Equal(t, result, expectedResult, "não foi ordenado corretamente")
}
