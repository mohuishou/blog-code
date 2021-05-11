package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_buildTree(t *testing.T) {
	tests := []struct {
		name     string
		preorder []int
		inorder  []int
		want     []int
	}{
		{
			preorder: []int{},
			inorder:  []int{},
			want:     nil,
		},
		{
			preorder: []int{3, 9, 20, 15, 7},
			inorder:  []int{9, 3, 15, 20, 7},
			want:     []int{3, 9, 20, null, null, 15, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildTree(tt.preorder, tt.inorder)
			assert.Equal(t, got, NewTreeNode(tt.want, 0))
		})
	}
}
