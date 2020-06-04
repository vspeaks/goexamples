package main

import "fmt"

func main() {
	var arr = []int32{11218, 100, 10,8,11, 7, 2, 1, 4, 3, 3, 0, 55, 6}
	merge_sort(arr)
}

func merge_sort(arr []int32) {
	merge_sortHelper(arr, 0, len(arr)-1)
}

func merge_sortHelper(arr []int32, start int, end int) []int32 {
	var mid int

	if start >= end {
		var retarr []int32
		retarr = append(retarr, arr[start])
		fmt.Printf("%v\n", retarr)
		return retarr
	}

	mid = (start + end) / 2
        return merge (  merge_sortHelper(arr, start, mid) , merge_sortHelper(arr, mid+1, end))
}

        func merge ( left []int32, right []int32)  []int32 {
	var  i, j int
	i = 0 
	j = 0
	var res []int32

	for {
         if i >= len(left) { break }
         if j >= len(right) { break }
	if left[i] <= right[j] {
                        res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}

	for {
		if i >= len(left) {
			break
		}
		res = append(res, left [i])
		i++
	}

	for {
		if j >= len(right) {
			break
		}
		res = append(res,right[j])
		j++
	}
	fmt.Printf("%v\n", res)
	return res
}
