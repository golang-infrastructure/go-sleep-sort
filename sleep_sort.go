package sleep_sort

import (
	"github.com/golang-infrastructure/go-gtypes"
	"sync"
	"time"
)

// SortUInt 对无符号序列排序，通过休眠
func SortUInt[T gtypes.Unsigned](slice []T) []T {
	if len(slice) == 0 {
		return slice
	}
	// 通过找到最小值，然后把所有值约掉一个最小值来提高效率
	min, _ := findMinAndMax(slice)
	resultChannel := make(chan T, len(slice))
	var wg sync.WaitGroup
	for _, x := range slice {
		localValue := x
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Second * time.Duration(localValue-min))
			resultChannel <- localValue
		}()
	}
	wg.Wait()
	close(resultChannel)
	resultSlice := make([]T, len(slice))
	for index := range resultSlice {
		resultSlice[index] = <-resultChannel
	}
	return resultSlice
}

// SortTime TODO 对时间排序
func SortTime(slice []time.Time) []time.Time {
	min, _ := findMinAndMaxByCompareFunc(slice, func(a, b time.Time) int {
		if a.Equal(b) {
			return 0
		} else if a.Before(b) {
			return -1
		} else {
			return 1
		}
	})
	resultChannel := make(chan time.Time, len(slice))
	var wg sync.WaitGroup
	for _, x := range slice {
		localValue := x
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Millisecond * time.Duration(localValue.UnixMilli()-min.UnixMilli()))
			resultChannel <- localValue
		}()
	}
	wg.Wait()
	close(resultChannel)
	resultSlice := make([]time.Time, len(slice))
	for index := range resultSlice {
		resultSlice[index] = <-resultChannel
	}
	return resultSlice
}

// 找出最大值和最小值
func findMinAndMax[T gtypes.Ordered](slice []T) (min T, max T) {
	return findMinAndMaxByCompareFunc(slice, func(a, b T) int {
		if a == b {
			return 0
		} else if a < b {
			return -1
		} else {
			return 1
		}
	})
}

// 找出最大值和最小值
func findMinAndMaxByCompareFunc[T any](slice []T, compareFunc func(a, b T) int) (min, max T) {
	if len(slice) == 0 {
		var zero T
		return zero, zero
	}
	min = slice[0]
	max = slice[0]
	for index := 1; index < len(slice); index++ {
		if compareFunc(slice[index], min) < 0 {
			min = slice[index]
		}
		if compareFunc(slice[index], max) > 0 {
			max = slice[index]
		}
	}
	return
}
