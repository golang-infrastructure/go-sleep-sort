package sleep_sort

import (
	"github.com/golang-infrastructure/go-gtypes"
	"time"
)

// SortUInt 对无符号序列排序，通过休眠
func SortUInt[T gtypes.Unsigned](slice []T) []T {
	resultChannel := make(chan T, len(slice))
	for _, x := range slice {
		localValue := x
		go func() {
			time.Sleep(time.Millisecond * time.Duration(localValue))
			resultChannel <- localValue
		}()
	}
	resultSlice := make([]T, len(slice))
	for index := range resultSlice {
		resultSlice[index] = <-resultChannel
	}
	return resultSlice
}

// SortTime TODO 对时间排序
func SortTime() {

}
