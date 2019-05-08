package main

import (
	"fmt"
	"os"
	"time"

	"randnum"
)

const (
	displayLen = 10
)

func main() {
	//load data
	data10 := randnum.GetData10()
	data1k := randnum.GetData1k()
	data5k := randnum.GetData5k()
	data10k := randnum.GetData10k()
	data50k := randnum.GetData50k()
	data100k := randnum.GetData100k()
	data200k := randnum.GetData200k()
	data500k := randnum.GetData500k()
	data1000k := randnum.GetData1000k()

	dataMap := map[string][]int32{
		"10":    data10,
		"1k":    data1k,
		"5k":    data5k,
		"10k":   data10k,
		"50k":   data50k,
		"100k":  data100k,
		"200k":  data200k,
		"500k":  data500k,
		"1000k": data1000k,
	}

	fmt.Println(len(data1k), len(data5k), len(data10k), len(data50k), len(data100k), len(data200k), len(data500k), len(data1000k))

	data := dataMap[os.Args[1]]

	data1 := make([]int32, len(data), len(data))
	copy(data1, data)
	//BubbleSort(data1)
	fmt.Println(data1[:displayLen], "\n")

	copy(data1, data)
	//SelectionSort(data1)
	fmt.Println(data1[:displayLen], "\n")

	copy(data1, data)
	//DirectInsertSort(data1)
	fmt.Println(data1[:displayLen], "\n")

	copy(data1, data)
	ShellSort(data1)
	fmt.Println(data1[:displayLen], "\n")

	copy(data1, data)
	start := time.Now().UnixNano()
	res := MergeSort(data1)
	end := time.Now().UnixNano()
	fmt.Println("MergeSort", "cmpCnt", 0, "swapCnt", 0, "elaped time:", "[ ", (end-start)/1e6, "ms", " ]")
	fmt.Println(res[:displayLen], "\n")

	copy(data1, data)
	start = time.Now().UnixNano()
	QuickSort1(data1)
	end = time.Now().UnixNano()
	fmt.Println("QuickSort1", "cmpCnt", 0, "swapCnt", 0, "elaped time:", "[ ", (end-start)/1e6, "ms", " ]")
	fmt.Println(data1[:displayLen], "\n")

	copy(data1, data)
	start = time.Now().UnixNano()
	QuickSort2(data1)
	end = time.Now().UnixNano()
	fmt.Println("QuickSort2", "cmpCnt", 0, "swapCnt", 0, "elaped time:", "[ ", (end-start)/1e6, "ms", " ]")
	fmt.Println(data1[:displayLen], "\n")

	copy(data1, data)
	start = time.Now().UnixNano()
	QuickSort3(data1)
	end = time.Now().UnixNano()
	fmt.Println("QuickSort3", "cmpCnt", 0, "swapCnt", 0, "elaped time:", "[ ", (end-start)/1e6, "ms", " ]")
	fmt.Println(data1[:displayLen], "\n")

	copy(data1, data)
	start = time.Now().UnixNano()
	QuickSort4(data1)
	end = time.Now().UnixNano()
	fmt.Println("QuickSort4", "cmpCnt", 0, "swapCnt", 0, "elaped time:", "[ ", (end-start)/1e6, "ms", " ]")
	fmt.Println(data1[:displayLen], "\n")

}

func BubbleSort(data []int32) {
	start := time.Now().UnixNano()
	var cmpCnt, swapCnt uint64

	for i := 1; i < len(data); i++ {
		for j := 0; j < len(data)-i; j++ {
			cmpCnt++
			if data[j] > data[j+1] {
				swapCnt++
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}

	end := time.Now().UnixNano()
	fmt.Println("BubbleSort", "cmpCnt", cmpCnt, "swapCnt", swapCnt, "elaped time:", "[ ", (end-start)/1e6, "ms", " ]")
}

func SelectionSort(data []int32) {
	start := time.Now().UnixNano()

	var cmpCnt, swapCnt uint64

	for i := 1; i < len(data); i++ {
		min := i - 1
		for j := min; j < len(data)-1; j++ {
			cmpCnt++
			if data[min] > data[j+1] {
				min = j + 1
			}
		}

		if min != i-1 {
			swapCnt++
			data[i-1], data[min] = data[min], data[i-1]
		}
	}
	end := time.Now().UnixNano()
	fmt.Println("SelectionSort", "cmpCnt", cmpCnt, "swapCnt", swapCnt, "elaped time:", "[ ", (end-start)/1e6, "ms", " ]")
}

func DirectInsertSort(data []int32) {
	start := time.Now().UnixNano()

	var cmpCnt, swapCnt uint64

	for i := 1; i < len(data); i++ {
		tmp := data[i]

		j := i - 1
		//少统计了比较次数
		for ; j >= 0 && data[j] > tmp; j-- {
			cmpCnt++
			swapCnt++
			data[j+1] = data[j]
		}

		if i != j+1 {
			swapCnt++
			data[j+1] = tmp
		}

	}
	end := time.Now().UnixNano()
	fmt.Println("DirectInsertSort", "cmpCnt", cmpCnt, "swapCnt", swapCnt, "elaped time:", "[ ", (end-start)/1e6, "ms", " ]")

}

func ShellSort(data []int32) {
	start := time.Now().UnixNano()
	var cmpCnt, swapCnt uint64

	//循环排列不同增量下的子序列
	for inc := len(data) / 2; inc >= 1; inc = inc / 2 {
		//循环排列每个子序列
		for i := 0; i < inc; i++ {
			//下面连个for循环形成标准的插入排序，只是相邻数据间隔为inc
			for j := i + inc; j < len(data); j = j + inc {
				tmp := data[j]
				k := j - inc

				for ; k >= 0 && data[k] > tmp; k = k - inc {
					cmpCnt++
					swapCnt++
					data[k+inc] = data[k]
				}

				if j != k+inc {
					swapCnt++
					data[k+inc] = tmp
				}
			}
		}
	}

	end := time.Now().UnixNano()
	fmt.Println("ShellInsertSort", "cmpCnt", cmpCnt, "swapCnt", swapCnt, "elaped time:", "[ ", (end-start)/1e6, "ms", " ]")
}

func merge(left, right []int32) (result []int32) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}

	if r == len(right) {
		result = append(result, left[l:]...)
	}
	if l == len(left) {
		result = append(result, right[r:]...)
	}
	return
}
func MergeSort(data []int32) (result []int32) {
	if len(data) <= 1 {
		return data
	}

	num := len(data) / 2

	left := MergeSort(data[:num])
	right := MergeSort(data[num:])

	result = merge(left, right)

	return
}

//用于找轴值，使得数列的分割更均衡。找到第一个值，中间位置一个值，最后一个值 三个值的中间值
func findPivot(data []int32) int {
	start := data[0]
	end := data[len(data)-1]
	mid := data[len(data)/2]

	if start > mid {
		if mid > end {
			return len(data) / 2
		}

		if start > end {
			return len(data) - 1
		} else {
			return 0
		}
	} else {
		if start > end {
			return 0
		}

		if mid < end {
			return len(data) / 2
		} else {
			return len(data) - 1
		}
	}
}

//对应网上的左右指针法，优化了选轴值
//所有的分割方法返回值都是轴值（分割点）的位置
//第一种分割方法采用首尾两指针向中间扫描的方法，指针相遇时停止
func partition1(data []int32) int32 {
	p := findPivot(data)

	//将轴值和尾部元素交换，将轴值存放在尾部
	if p != len(data)-1 {
		data[p], data[len(data)-1] = data[len(data)-1], data[p]
		p = len(data) - 1
	}

	var i, j int32
	j = int32(p - 1)

	for i < j {
		//j从后往前走，直到找到一个小于轴值的元素停下来
		for i < j && data[i] <= data[p] {
			i++
		}

		for i < j && data[j] >= data[p] {
			j--
		}

		data[i], data[j] = data[j], data[i]
	}

	if data[i] > data[p] {
		data[i], data[p] = data[p], data[i]
	} else {
		return int32(p)
	}
	//fmt.Println("i:", data[i], "p:", data[p])

	return i
}

//对应网上的左右指针法，简单选轴值
//选取最后一个元素为轴值
func partition2(data []int32) int32 {
	var i, j, p int32
	p = int32(len(data) - 1)
	j = p - 1

	for i < j {
		for i < j && data[i] <= data[p] {
			i++
		}

		for i < j && data[j] >= data[p] {
			j--
		}

		data[i], data[j] = data[j], data[i]
	}

	if data[i] > data[p] {
		data[i], data[p] = data[p], data[i]
	} else {
		return p
	}

	return i
}

//此方法对应于网上的前后指针法
//选取最后一个为轴值，两个下标均从左边向右移动
//j在前面，遇到大于等于轴值的继续前进，直到找到小于轴值的停下来
//此时i往前走一步，交换i和j的值
func partition3(data []int32) int32 {
	var i, j, p int32
	p = int32(len(data) - 1)
	i = -1
	j = 0

	for j < p {
		if data[j] < data[p] {
			i++
			data[i], data[j] = data[j], data[i]
		}

		j++
	}

	data[i+1], data[j] = data[j], data[i+1]

	return i + 1
}

//此方法对应于网上的挖坑法
//选择第一个为轴值，两个下标从两端向中间移动，相遇则结束
func partition4(data []int32) int32 {
	var i, j, p, pvalue int32
	p = 0
	i = 0
	j = int32(len(data) - 1)
	//pvalue保存轴值本身
	pvalue = data[0]

	for i < j {
		for i < j && data[j] >= pvalue {
			j--
		}

		if i == j {
			break
		}

		data[p] = data[j]
		p = j

		for i < j && data[i] <= pvalue {
			i++
		}

		if i < j {
			data[p] = data[i]
			p = i
		}
	}

	data[p] = pvalue

	return p
}

func QuickSort1(data []int32) {
	if len(data) < 2 {
		return
	}

	p := partition1(data)

	QuickSort1(data[:p])
	QuickSort1(data[p+1:])
}
func QuickSort2(data []int32) {
	if len(data) < 2 {
		return
	}

	p := partition2(data)

	QuickSort2(data[:p])
	QuickSort2(data[p+1:])
}

func QuickSort3(data []int32) {
	if len(data) < 2 {
		return
	}

	p := partition3(data)

	QuickSort3(data[:p])
	QuickSort3(data[p+1:])
}

func QuickSort4(data []int32) {
	if len(data) < 2 {
		return
	}

	p := partition4(data)

	QuickSort4(data[:p])
	QuickSort4(data[p+1:])
}

/*
//二分查找插入容易造成排序不稳定，需要逻辑上保证稳定
func BinaryInsertSort(data []int32) {
	start := time.Now().UnixNano()

	var cmpCnt, swapCnt uint64

	for i := 1; i < len(data); i++ {
		tmp := data[i]

		j := i - 1

		if data[j] <= tmp{
			continue
		}

		start := 0
		end := j
		mid := (start + end) / 2
		for ; mid > start; mid = (start + end) / 2 {
			cmpCnt++

			if data[mid] > tmp {
				end = mid
			} else if data[mid] < tmp {
				start = mid
			} else {
				k := mid + 1
				for ; k <= j; k++ {
					cmpCnt++
					if data[k] > tmp {
						break
					}
				}

				for ; j >= k; j-- {
					swapCnt++
					data[j+1] = data[j]
				}
				swapCnt++
				data[k] = tmp
				break
			}
		}

		k := mid + 1
		for ; j >= k; j-- {
			swapCnt++
			data[j+1] = data[j]
		}
		swapCnt++
		data[k] = tmp
	}
	end := time.Now().UnixNano()
	fmt.Println("BinaryInsertSort", "cmpCnt", cmpCnt, "swapCnt", swapCnt, "elaped time:", "[ ", (end-start)/1e6, "ms", " ]")

}*/
