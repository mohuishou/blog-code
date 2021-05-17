package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isCompleteTree(t *testing.T) {
	tests := []struct {
		name string
		tree []int
		want bool
	}{
		{
			tree: []int{1, 2, 3, 4, 5, 6},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isCompleteTree(NewTreeNode(tt.tree, 0)))
		})
	}
}
