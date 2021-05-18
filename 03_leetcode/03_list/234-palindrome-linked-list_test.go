package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name string
		list []int
		want bool
	}{
		{
			list: []int{1, 2, 2, 1},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isPalindrome(NewList(tt.list)))
		})
	}
}
