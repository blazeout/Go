package main

import (
	"fmt"
	"math/rand"
)

func insertSort(arr []int) {
	// 将数组分为已经排序好的区间[0, bound), 和未排序好的区间[bound,length]
	for i := 0; i < len(arr); i++ {
		end := i-1
		cur := arr[i]
		// 如果遇到了比cur大的元素就需要往后搬
		for end >= 0 && arr[end] > cur  {
			arr[end+1] = arr[end]
			end--
		}
		arr[end+1] = cur
	}
}

func insertSortWithGap(arr []int, gap int) {
	for i := 0; i < len(arr); i++ {
		end := i-gap
		cur := arr[i]
		for end >= 0 && arr[end] > cur {
			arr[end+gap] = arr[end]
			end -= gap
		}
		arr[end+gap] = cur
	}
}

func shellSort(arr []int) {
	// 如果会插入排序那么希尔排序也就会, 希尔排序就是将插入排序分堆利用gap来分组块
	// 常用的gap就是长度依次除2
	gap := len(arr)/2
	for gap > 0 {
		insertSortWithGap(arr, gap)
		gap  /= 2
	}
}

func selectSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	// 每次定义一个position, 用来找到后面的数组中的最小值
	for i := 0; i < len(arr)-1; i++ {
		// 以本次的i为下标先当做最小的元素
		position := i
		for j := position+1; j < len(arr); j++ {
			if arr[j] < arr[position] {
				position = j
			}
		}
		arr[i], arr[position] = arr[position], arr[i]
	}
}

func heapSort(arr []int) {
	size := len(arr)
	createHeap(arr, size)
	for i := len(arr)-1 ; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		size--
		shiftDown(arr, 0, size)
	}
}

func createHeap(arr []int, size int) {
	// 堆排序的核心思想就是, 先建立一个大跟堆, 再依次将数组第一个元素和最后一个元素交换, 直至堆的大小0
	// 从第一个非叶子结点开始建堆
	root := (len(arr)-1-1) >> 1 // size == len(arr) size-1是最后一个元素的下标, 那么他的父亲结点就是 parent = (size-1-1)/2
	for ; root >= 0; root-- {
		// 依次往上调整
		shiftDown(arr, root, size)
	}
}

func shiftDown(arr []int, parent int, size int) {
	// 首先找到左孩子
	child := parent*2 + 1
	for child < size {
		// 一直向下调整
		// 判断左子树和右子数哪个大
		if child+1 < size && arr[child] < arr[child+1] {
			child += 1
		}
		// 判断父亲的值和孩子的值哪个大, 大了的话 就需要交换
		if arr[parent] < arr[child] {
			arr[parent], arr[child] = arr[child], arr[parent]
		}else {
			// 不需要交换直接break
			break
		}
		parent = child
		child = 2*parent +1
	}
}

func quickSort(nums []int, l,r int) {
	if l >= r { return }
	pivot := partition(nums, l, r)
	quickSort(nums, l, pivot-1)
	quickSort(nums, pivot + 1, r)
	return
}

func partition(nums []int, l,r int) int {
	pivot := l + rand.Intn(r-l+1)
	nums[pivot], nums[r] = nums[r], nums[pivot]
	left := l-1
	right := l
	for ; right < r; right++ {
		if nums[right] <= nums[r] {
			left++
			nums[left], nums[right] = nums[right], nums[left]
		}
	}
	nums[right] ,nums[left+1] = nums[left+1], nums[right]
	return left+1
}


func merge(left, right []int) []int {
	tmp := make([]int, 0)
	i, j := 0,0
	l, r := len(left), len(right)
	for i < l && j < r {
		if left[i] > right[j] {
			tmp = append(tmp, right[j])
			j++
		}else if left[i] <= right[j] {
			tmp = append(tmp, left[i])
			i++
		}
	}
	tmp = append(tmp, left[i:]...)
	tmp = append(tmp, right[j:]...)
	return tmp
}

func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	// 分治
	i := len(nums)/2
	left := mergeSort(nums[:i])
	right := mergeSort(nums[i:])
	result := merge(left, right)
	return result
}

func main() {
	var arr = []int{2,4,1,5,6,3,8}

	fmt.Println(arr)
	//insertSort(arr)
	//shellSort(arr)
	//selectSort(arr)
	//heapSort(arr)
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
