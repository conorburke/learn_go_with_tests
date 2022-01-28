package slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
    t.Run("sums up ints in slice", func(t *testing.T) {
        nums := []int{1,2,3,4,5,6,7,8,9,10}
        expected := 55
        got := Sum(nums)

        if expected != got {
            t.Errorf("expected '%d' but got '%d'", expected, got)
        }
    })
}

func TestSumAll(t *testing.T) {
    t.Run("tests one or more input slices", func(t *testing.T) {
        nums1 := []int{1,2,3}
        nums2 := []int{1,2,3,4}

        expected := []int{6, 10}
        got := SumAll(nums1, nums2)

        // if expected != got {
        if !reflect.DeepEqual(expected, got) {
            t.Errorf("expected '%d' but got '%d'", expected, got)
        }
    })
}