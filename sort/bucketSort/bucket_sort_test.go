package bucketSort

import (
	"testing"
)

func TestBucketSort(t *testing.T) {
	a := []int{1, 2, 1, 5, 6, 2, 5, 1, 3, 7, 3, 4, 4, 6, 7, 6, 6, 5, 3}
	BucketSort(a)
	t.Log(a)
}
