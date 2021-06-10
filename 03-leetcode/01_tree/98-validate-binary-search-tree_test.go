package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isValidBST(t *testing.T) {

	tests := []struct {
		name string
		tree []int
		want bool
	}{
		{
			tree: []int{1, 1},
			want: false,
		},
		{
			tree: []int{5, 4, 6, null, null, 3, 7},
			want: false,
		},
		{
			tree: []int{5, 1, 4, null, null, 3, 6},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isValidBST(NewTreeNode(tt.tree, 0)))
		})
	}
}
