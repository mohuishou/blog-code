package list

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getIntersectionNode(t *testing.T) {
	headA := NewList([]int{4, 1, 8, 4, 5})
	headB := NewList([]int{5, 6, 1})
	headB.Next.Next.Next = headA.Next.Next
	// headA := NewList([]int{2, 6, 4})
	// headB := NewList([]int{5, 1})
	fmt.Println(headA.array(), headB.array())
	assert.Equal(t, headA.Next.Next, getIntersectionNode(headA, headB))
}
