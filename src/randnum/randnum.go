package randnum

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	Data10Path    = "data/num10.txt"
	Data1kPath    = "data/num1k.txt"
	Data5kPath    = "data/num5k.txt"
	Data10kPath   = "data/num10k.txt"
	Data50kPath   = "data/num50k.txt"
	Data100kPath  = "data/num100k.txt"
	Data200kPath  = "data/num200k.txt"
	Data500kPath  = "data/num500k.txt"
	Data1000kPath = "data/num1000k.txt"
	Data10        = 10
	Data1k        = 1000
	Data5k        = 5000
	Data10k       = 10000
	Data50k       = 50000
	Data100k      = 100000
	Data200k      = 200000
	Data500k      = 500000
	Data1000k     = 1000000
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func RandGenerator(n uint32, path string) {
	rand.Seed(time.Now().UnixNano())

	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file fail.", err)
		return
	}
	defer file.Close()

	w := bufio.NewWriter(file)

	var i uint32
	for ; i < n; i++ {
		_, err := fmt.Fprintln(w, rand.Intn(100000))
		if err != nil {
			fmt.Println("write fail", err)
		}
	}

	w.Flush()
}
func GetData10() []int32 {
	data := make([]int32, 0, Data10)

	if getData(&data, Data10Path, Data10) {
		return data
	}

	fmt.Println("GetData10 fail")

	return nil
}
func GetData1k() []int32 {
	data := make([]int32, 0, Data1k)

	if getData(&data, Data1kPath, Data1k) {
		return data
	}

	fmt.Println("GetData1k fail")

	return nil
}

func GetData5k() []int32 {
	data := make([]int32, 0, Data5k)

	if getData(&data, Data5kPath, Data5k) {
		return data
	}

	fmt.Println("GetData5k fail")

	return nil
}

func GetData10k() []int32 {
	data := make([]int32, 0, Data10k)

	if getData(&data, Data10kPath, Data10k) {
		return data
	}

	fmt.Println("GetData10k fail")

	return nil
}

func GetData50k() []int32 {
	data := make([]int32, 0, Data50k)

	if getData(&data, Data50kPath, Data50k) {
		return data
	}

	fmt.Println("GetData50k fail")

	return nil
}

func GetData100k() []int32 {
	data := make([]int32, 0, Data100k)

	if getData(&data, Data100kPath, Data100k) {
		return data
	}

	fmt.Println("GetData100k fail")

	return nil
}

func GetData200k() []int32 {
	data := make([]int32, 0, Data200k)

	if getData(&data, Data200kPath, Data200k) {
		return data
	}

	fmt.Println("GetData200k fail")

	return nil
}

func GetData500k() []int32 {
	data := make([]int32, 0, Data500k)

	if getData(&data, Data500kPath, Data500k) {
		return data
	}

	fmt.Println("GetData500k fail")

	return nil
}

func GetData1000k() []int32 {
	data := make([]int32, 0, Data1000k)

	if getData(&data, Data1000kPath, Data1000k) {
		return data
	}

	fmt.Println("GetData1000k fail")

	return nil
}

func getData(data *[]int32, path string, num uint32) bool {
	b, err := PathExists(path)
	if err == nil {
		if b == false {
			RandGenerator(num, path)
			fmt.Println("create new data file", path)
		}
	}

	f, err := os.Open(path)
	if err != nil {
		fmt.Println("open file fail", path)
		return false
	}

	defer f.Close()

	rd := bufio.NewReader(f)

	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				if line != "" {
					fmt.Println("no newline character at the end of last line")
				} else {
					//fmt.Println("end of file")
					break
				}
			} else {
				break
			}
		}

		num, err := strconv.ParseInt(strings.TrimSpace(line), 10, 32)
		if err != nil {
			fmt.Println("fail to convert string")
			continue
		}

		*data = append(*data, int32(num))
	}

	return true
}
