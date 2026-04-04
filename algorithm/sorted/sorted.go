package sorted

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
