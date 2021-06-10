package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_wordBreak(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "",
			args: args{
				s:        "catsanddog",
				wordDict: []string{"cat", "cats", "and", "sand", "dog"},
			},
			want: []string{
				"cat sand dog",
				"cats and dog",
			},
		},
		{
			name: "",
			args: args{
				s:        "catsandog",
				wordDict: []string{"cat", "cats", "and", "sand", "dog"},
			},
			want: []string{},
		},
		{
			name: "",
			args: args{
				s:        "pineapplepenapple",
				wordDict: []string{"apple", "pen", "applepen", "pine", "pineapple"},
			},
			want: []string{
				"pine applepen apple",
				"pineapple pen apple",
				"pine apple pen apple",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := wordBreak(tt.args.s, tt.args.wordDict)
			assert.Equal(t, tt.want, got)
		})
	}
}
