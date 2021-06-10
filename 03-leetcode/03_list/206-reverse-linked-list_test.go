package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseList(t *testing.T) {

	tests := []struct {
		name string
		list []int
		want []int
	}{
		{
			list: []int{1, 2, 3, 4, 5},
			want: []int{5, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, reverseList(NewList(tt.list)).array())
		})
	}
}
