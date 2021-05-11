package tree

import (
	"reflect"
	"testing"
)

func Test_rightSideView(t *testing.T) {

	tests := []struct {
		name string
		tree []int
		want []int
	}{
		{
			tree: []int{1, 2, 3, null, 5, null, 4},
			want: []int{1, 3, 4},
		},
		{
			tree: []int{},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := NewTreeNode(tt.tree, 0)
			if got := rightSideView(root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rightSideView() = %v, want %v", got, tt.want)
			}
		})
	}
}
