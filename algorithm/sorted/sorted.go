package sorted

import "math/rand"

// SelectionSort 选择排序
func SelectionSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		minIndex := i
		for j := i + 1; j < length; j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}

		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}

// InsertionSort 插入排序
func InsertionSort(arr []int) []int {
	length := len(arr)

	for i := 0; i < length; i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			} else {
				break
			}
		}
	}

	return arr
}

// MergeSort 归并排序
func MergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	var mergeSort func(arr []int, l, r int)

	// 临时数组 优化空间
	tmp := make([]int, len(arr))

	mergeSort = func(arr []int, l, r int) {
		if l >= r {
			return
		}

		mid := l + (r-l)/2

		// 左半部分递归
		mergeSort(arr, l, mid)

		// 右半部分递归
		mergeSort(arr, mid+1, r)

		// 两半部分有序合并
		if arr[mid] > arr[mid+1] {
			merge(arr, tmp, l, mid, r)
		}

	}

	mergeSort(arr, 0, len(arr)-1)
}

// MergeSortBU 自底向上归并排序
func MergeSortBU(arr []int) {
	if len(arr) <= 1 {
		return
	}

	// 临时数组 优化空间
	tmp := make([]int, len(arr))

	// 遍历合并的两个区间的起始位置i
	// 合并[i, i+sz-1] 和[i+sz, i+sz+sz-1] min(i+sz+sz-1, len(arr)-1)
	for sz := 1; sz < len(arr); sz += sz {
		for i := 0; i+sz < len(arr); i += sz + sz {
			merge(arr, tmp, i, i+sz-1, min(i+sz+sz-1, len(arr)-1))
		}
	}
}

func merge(arr, tmp []int, l, mid, r int) {
	idx := l
	i, j := l, mid+1

	for i <= mid && j <= r {
		if arr[i] <= arr[j] {
			tmp[idx] = arr[i]

			i++
		} else {
			tmp[idx] = arr[j]
			j++
		}
		idx++
	}
	for i <= mid {
		tmp[idx] = arr[i]
		idx++
		i++
	}

	for j <= r {
		tmp[idx] = arr[j]
		idx++
		j++
	}

	copy(arr[l:r+1], tmp[l:r+1])
}

// QuickSort 快速排序
func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	var quickSort func(arr []int, l, r int)

	quickSort = func(arr []int, l, r int) {
		if l >= r {
			return
		}

		p := partition2(arr, l, r)

		// 左半部分
		quickSort(arr, l, p-1)

		// 右半部分
		quickSort(arr, p+1, r)
	}

	quickSort(arr, 0, len(arr)-1)
}

// [l,p-1] [p+1, r]
func partition(arr []int, l, r int) int {
	// [l+1, j] < v, [j + 1, i -1] > v
	j := l
	for i := l + 1; i <= r; i++ {
		if arr[i] < arr[l] {
			j++
			arr[j], arr[i] = arr[i], arr[j]
		}
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}

func partition2(arr []int, l, r int) int {
	// [l+1, j] < v, [j + 1, i -1] > v
	mid := l + rand.Intn(r-l+1)
	arr[mid], arr[l] = arr[l], arr[mid]
	j := l
	for i := l + 1; i <= r; i++ {
		if arr[i] < arr[l] {
			j++
			arr[j], arr[i] = arr[i], arr[j]
		}
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}

// 双路partition
func partition2ways(arr []int, l, r int) int {
	p := l + rand.Intn(r-l+1)
	arr[p], arr[l] = arr[l], arr[p]

	// [l+1, i-1] <= v   [j+1, r] >=v
	i, j := l+1, r

	for {
		for i <= j && arr[i] < arr[l] {
			i++
		}
		for i <= j && arr[j] > arr[l] {
			j--
		}
		if i >= j {
			break
		}

		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	arr[l], arr[j] = arr[j], arr[l]

	return j
}

// 三路快排
func quickSort3Ways(arr []int, l, r int) {
	if l >= r {
		return
	}

	lt, gt := partition3ways(arr, l, r)

	// 左半部分
	quickSort3Ways(arr, l, lt)

	// 右半部分
	quickSort3Ways(arr, gt, r)
}

// 三路partition
func partition3ways(arr []int, l, r int) (int, int) {

	p := l + rand.Intn(r-l+1)
	arr[p], arr[l] = arr[l], arr[p]

	// [l+1, lt] < v， [lt+1, i-1] == v， [gt, r] > v
	lt, gt, i := l, r+1, l+1
	for i < gt {
		if arr[i] < arr[l] {
			lt++
			arr[i], arr[lt] = arr[lt], arr[i]
			i++
		} else if arr[i] > arr[l] {
			gt--
			arr[i], arr[gt] = arr[gt], arr[i]
		} else { // arr[i] == arr[l]
			i++
		}
	}

	// [l+1, lt - 1] < v， [lt, i-1] == v， [gt, r] > v
	arr[lt], arr[l] = arr[l], arr[lt]
	lt++

	return lt - 1, gt
}

func HeapSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	// 整理成一个最大堆
	heapify(arr)

	for i := len(arr) - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		siftDown(arr, 0, i)
	}

}

func heapify(arr []int) {
	if len(arr) <= 1 {
		return
	}
	for i := (len(arr) - 2) / 2; i >= 0; i-- {
		siftDown(arr, i, len(arr))
	}
}

// 对[0, n) 中对索引为idx的进行siftDown
func siftDown(arr []int, idx, n int) {
	//leftIdx, rightIdx := 2* idx + 1, 2* idx + 2
	for 2*idx+1 < n {
		maxIndex := 2*idx + 1
		if maxIndex+1 < n && arr[maxIndex+1] > arr[maxIndex] {
			maxIndex++
		}

		if arr[idx] >= arr[maxIndex] {
			break
		}

		// 交换数据
		arr[idx], arr[maxIndex] = arr[maxIndex], arr[idx]
		// idx为最大的这个元素
		idx = maxIndex
	}
}
