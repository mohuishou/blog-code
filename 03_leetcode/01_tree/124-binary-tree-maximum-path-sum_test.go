package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxPathSum(t *testing.T) {
	tests := []struct {
		name string
		tree []int
		want int
	}{
		{
			tree: []int{1, 2, 3},
			want: 6,
		},
		{
			tree: []int{-10, 9, 20, null, null, 15, 7},
			want: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, maxPathSum(NewTreeNode(tt.tree, 0)))
		})
	}
}
