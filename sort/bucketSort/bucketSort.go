package bucketSort

import (
	"algo/sort/quickSort"
)

/**
  桶排序，顾名思义，会用到“桶”，核心思想是将要排序的数据分到几个有序的桶里，
  每个桶里的数据再单独进行排序。桶内排完序之后，再把每个桶里的数据按照顺序依次取出，
  组成的序列就是有序的了
*/
func BucketSort(a []int) {
	if a == nil {
		return
	}

	length := len(a)

	if length <= 1 {
		return
	}

	bucket := make([][]int, length)

	for _, value := range a {
		bucket[value] = append(bucket[value], value)
	}

	tmpPos := 0
	for _, row := range bucket {
		if len(row) > 1 {
			quickSort.QuickSort(row)
		}
		copy(a[tmpPos:], row)
		tmpPos += len(row)
	}

}
