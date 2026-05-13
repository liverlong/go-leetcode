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

// 13. 罗马数字转整数
func romanToInt(s string) int {
	roman := map[byte]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50,
		'C': 100, 'D': 500, 'M': 1000,
	}
	n := len(s)
	res := roman[s[n-1]]          // 1. 先取最后一个字符的值
	for i := n - 2; i >= 0; i-- { // 2. 从倒数第二个开始向左遍历
		if roman[s[i]] < roman[s[i+1]] {
			res -= roman[s[i]] // 3. 当前值 < 右边值 → 减法
		} else {
			res += roman[s[i]] // 4. 否则 → 加法
		}
	}
	return res
}
