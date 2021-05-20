package main

import (
	"fmt"
	"sort"
)

func main() {
	arr1 := []int{1, 2}
	arr2 := []int{3, 4}
	mArr := append(arr1, arr2...)
	sort.Ints(mArr)
	fmt.Printf("Sorted merged array: %v\n", mArr)
	med := findMedianSortedArrays(arr1, arr2)
	fmt.Printf("Median: %f\n", med)
}

// mergedArr = merge(nums1, nums2)
// l = len(mergedArr)
// sort(mergedArr)
// if l is odd
//   return (mergedArr[(l-1)/2] + mergedArr[(l+1)/2])/2
// else
//   return mergedArr[l/2]

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	mArr := append(nums1, nums2...) // merged array
	sort.Ints(mArr) // sort the arr
	l := len(mArr)
	if l == 0 {
		return float64(mArr[0])
	}
	if l%2 == 0 {
		return float64(float64(mArr[(l-1)/2] + mArr[(l+1)/2])/float64(2))
	} else {
		return float64(mArr[l/2])
	}
}
