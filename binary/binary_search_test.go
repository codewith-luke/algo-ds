package binary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	arg1     int
	expected int
}

func TestSearch(t *testing.T) {
	assert := assert.New(t)
	list := []int{10, 20, 30, 40, 50}

	tests := []testCase{
		{10, 0},
		{50, 4},
		{20, 1},
		{30, 2},
		{40, 3},
	}

	_, err := Search(list, 100)
	assert.Equal(err.Error(), "none", "should return none error")

	for _, test := range tests {
		pos, _ := Search(list, test.arg1)
		assert.Equal(pos, test.expected, "should return correct position")
	}
}
