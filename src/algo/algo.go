package main

import (
	"fmt"
	"os"
	"time"

	"randnum"
)

func main() {
	//load data
	data1k := randnum.GetData1k()
	data5k := randnum.GetData5k()
	data10k := randnum.GetData10k()
	data50k := randnum.GetData50k()
	data100k := randnum.GetData100k()
	data200k := randnum.GetData200k()
	data500k := randnum.GetData500k()
	data1000k := randnum.GetData1000k()

	dataMap := map[string][]int32{
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
	fmt.Println(data1[:50], "\n")

	copy(data1, data)
	//SelectionSort(data1)
	fmt.Println(data1[:50], "\n")

	copy(data1, data)
	//DirectInsertSort(data1)
	fmt.Println(data1[:50], "\n")

	copy(data1, data)
	ShellSort(data1)
	fmt.Println(data1[:50], "\n")

	copy(data1, data)
	start := time.Now().UnixNano()
	res := MergeSort(data1)
	end := time.Now().UnixNano()
	fmt.Println("MergeSort", "cmpCnt", 0, "swapCnt", 0, "elaped time:", "[ ", (end-start)/1e6, "ms", " ]")
	fmt.Println(res[:50], "\n")
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
