package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_widthOfBinaryTree(t *testing.T) {

	tests := []struct {
		name string
		tree []int
		want int
	}{
		{
			tree: []int{1, 3, 2, 5, null, null, 9, 6, null, null, 7},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, widthOfBinaryTree(NewTree(tt.tree)))
		})
	}
}
