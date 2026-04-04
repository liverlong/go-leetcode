package easy

// 有效的括号
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	res := make([]rune, 0)

	for _, c := range s {
		if c == '{' || c == '[' || c == '(' {
			res = append(res, c)
		} else {
			if len(res) == 0 {
				return false
			}

			if c == '}' && res[len(res)-1] != '{' {
				return false
			}

			if c == ']' && res[len(res)-1] != '[' {
				return false
			}

			if c == ')' && res[len(res)-1] != '(' {
				return false
			}

			res = res[:len(res)-1]
		}
	}

	return len(res) == 0
}

// 206 反转链表
func reverseList(head *ListNode) *ListNode {

	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next

		// reverse
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func reverseList2(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}
	rev := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil

	return rev
}
