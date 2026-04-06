package structure

// 最大堆：根节点的元素大于等于所有子树中节点的值；最小堆：根节点的元素小于等于所有子树中节点的值

// i 从0开始： parent(i) = i-1 / 2；left child(i) = 2 * i + 1；right child = 2 * i + 2

// MaxHeap 最大堆
type MaxHeap struct {
	data []int
}

func NewMaxHeap(arr []int) *MaxHeap {
	maxHeap := &MaxHeap{
		data: arr,
	}

	// heapify 从第一个非叶子节点
	for i := maxHeap.parent(len(maxHeap.data) - 1); i >= 0; i-- {
		maxHeap.siftDown(i)
	}

	return maxHeap
}

// Size 堆大小
func (m *MaxHeap) Size() int {
	return len(m.data)
}

// IsEmpty 堆是否为空
func (m *MaxHeap) IsEmpty() bool {
	return len(m.data) == 0
}

// Add 向堆中添加元素
func (m *MaxHeap) Add(val int) {
	m.data = append(m.data, val)
	m.siftUp(len(m.data) - 1)
}

func (m *MaxHeap) RemoveMax() int {
	if len(m.data) == 0 {
		return -1
	}

	ret := m.data[0]

	m.data[0], m.data[len(m.data)-1] = m.data[len(m.data)-1], m.data[0]
	m.data = m.data[:len(m.data)-1]

	m.siftDown(0)
	return ret
}

// Replace 堆中移除最大元素 添加新元素
func (m *MaxHeap) Replace(val int) int {
	ret := m.data[0]

	m.data[0] = val

	m.siftDown(0)

	return ret
}

// ～～～～～～～～～～～～～～～～～private method ~~~~~~~~~~~~~~~~~~
// idx 父节点的索引
func (m *MaxHeap) parent(idx int) int {
	return (idx - 1) / 2
}

// idx 左孩子的索引
func (m *MaxHeap) leftChild(idx int) int {
	return 2*idx + 1
}

// idx 右边孩子的索引
func (m *MaxHeap) rightChild(idx int) int {
	return 2*idx + 2
}

// 上浮到正确位置
func (m *MaxHeap) siftUp(idx int) {
	// 子节点值大于父节点的值 需要上浮
	for idx > 0 && m.data[m.parent(idx)] < m.data[idx] {
		m.data[m.parent(idx)], m.data[idx] = m.data[idx], m.data[m.parent(idx)]
		idx = m.parent(idx)
	}
}

// 下沉到最大位置
func (m *MaxHeap) siftDown(idx int) {
	// 循环条件 是否到叶子节点
	for m.leftChild(idx) < len(m.data) {
		maxIdx := m.leftChild(idx)

		// 右节点存在且大于左节点的数
		if maxIdx+1 < len(m.data) && m.data[maxIdx+1] > m.data[maxIdx] {
			maxIdx++
		}

		// 父节点最大
		if m.data[idx] >= m.data[maxIdx] {
			break

		}

		// 交换父节点和最大值节点的值 并且更新idx
		m.data[idx], m.data[maxIdx] = m.data[maxIdx], m.data[idx]
		idx = maxIdx
	}
}
