package generateParentheses

import (
	"reflect"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	tests := []struct {
		name  string
		count int
		want  []string
	}{
		{
			"one",
			1,
			[]string{"()"},
		},
		{
			"three",
			3,
			[]string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateParenthesis(tt.count); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
			}
		})
	}
}
