package structure

type node struct {
	isWord bool
	next   map[rune]*node
}

// Trie 前缀树
type Trie struct {
	root *node
	size int
}

func Constructor() *Trie {
	return &Trie{
		root: &node{next: make(map[rune]*node)},
	}
}

// Add 前缀树中加入字符串
func (t *Trie) Add(word string) {

	cur := t.root
	for _, c := range word {
		// 不存在 则添加一个
		if _, ok := cur.next[c]; !ok {
			cur.next[c] = &node{next: make(map[rune]*node)}
		}
		cur = cur.next[c]
	}
	// 原来就不是个新单词
	if !cur.isWord {
		cur.isWord = true
		t.size++
	}

}

// Contains 前缀树中是否包含word
func (t *Trie) Contains(word string) bool {
	cur := t.root

	for _, c := range word {
		if _, ok := cur.next[c]; !ok {
			return false
		}
		cur = cur.next[c]
	}

	return cur.isWord

}

func (t *Trie) StartsWith(prefix string) bool {
	cur := t.root

	for _, c := range prefix {
		if _, ok := cur.next[c]; !ok {
			return false
		}
		cur = cur.next[c]
	}

	return true
}

func (t *Trie) Delete(word string) bool {
	if word == "" {
		return false
	}

	deleted := t.delete(t.root, []rune(word), 0)
	if deleted {
		t.size--
	}

	return deleted
}

func (t *Trie) delete(cur *node, word []rune, index int) bool {

	// 终止条件 到了单词末尾
	if index == len(word) {
		// 如果单词不存在 直接返回
		if !cur.isWord {
			return false
		}

		cur.isWord = false
		// 当前节点没有子节点 说明可以物理删除
		return len(cur.next) == 0

	}

	c := word[index]
	nextNode, ok := cur.next[c]
	if !ok {
		// 说明要删除的单词不在树中
		return false
	}

	// 递归向下找
	canDeleteChild := t.delete(nextNode, word, index+1)

	if canDeleteChild {
		// 在这个路径中删除c
		delete(cur.next, c)

		// 如果当前节点满足两个条件 也可以被删除
		// 1、当前节点不是另外一个单词的结尾 isWord = false
		// 2、当前节点已经没有其他子节点 len(cur.next) == 0

		return !cur.isWord && len(cur.next) == 0
	}

	return false

}
