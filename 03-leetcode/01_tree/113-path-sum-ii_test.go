package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_pathSum(t *testing.T) {
	type args struct {
		root      []int
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				root:      []int{5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1},
				targetSum: 22,
			},
			want: [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := NewTreeNode(tt.args.root, 0)
			t.Log(tree.array())
			assert.Equal(t, tt.want, pathSum(tree, tt.args.targetSum))
		})
	}
}
