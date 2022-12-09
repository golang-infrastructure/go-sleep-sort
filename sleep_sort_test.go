package sleep_sort

import (
	"testing"
)

func TestSortUInt(t *testing.T) {

	slice := []uint{8, 6, 9, 3}
	r := SortUInt(slice)
	t.Log(r)

	slice = []uint{1000003, 1000010, 1000000}
	r = SortUInt(slice)
	t.Log(r)

}

func TestSortTime(t *testing.T) {
	// TODO
}
