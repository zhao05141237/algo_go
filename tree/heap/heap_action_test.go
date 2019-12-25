package heap

import (
	"math/rand"
	"testing"
	"time"
)

func TestHeap_InsertData(t *testing.T) {
	//堆排序简单版本
	rand.Seed(time.Now().UnixNano())
	h := NewHeap(10)
	//建最大堆
	for i := 1; i <= 10; i++ {
		h.InsertData(rand.Intn(100), MaxTopSort)
	}
	h.printHeap()
	//每次删除顶元素 循环n-1次 从而形成从有序数组
	for i := 1; i < 10; i++ {
		h.RemoveTopData(MaxTopSort)
	}
	h.printHeap()
}

func TestMergedSortedData(t *testing.T) {
	//合并有序数组
	const MaxSize = 4
	a := [MaxSize][]int{
		{9, 10, 12, 30, 32, 32, 39, 61, 77, 87},
		{23, 27, 30, 63, 73, 72, 67, 79, 95, 97},
		{1, 5, 6, 18, 24, 56, 62, 65, 88, 92},
		{7, 15, 30, 31, 71, 56, 41, 43, 72, 97},
	}
	h := NewHeap(MaxSize + 1)
	b := make([]int, MaxSize)        //存放每个数组每次推进的位置
	countRow := make([]int, MaxSize) //存放每个数组的长度
	sumCountRow := 0
	var c []int //存放最终有序数组
	i := 0
	row := 0
	col := 0
	//首先用每个有序数组的首元素建最小堆
	for i < MaxSize {
		countRow[i] = len(a[i])
		sumCountRow += countRow[i]
		h.InsertData(a[i][0], MinTopSort)
		i++
	}
	//每次取堆顶元素放置最终排序数组
	for j := 0; j < sumCountRow; j++ {
		h.RemoveTopData(MinTopSort)   //取堆顶元素
		c = append(c, h.a[h.count+1]) //堆顶元素追加到最终数组
		for i = 0; i < MaxSize; i++ {
			//确定推进哪个有序数组
			if b[i] < countRow[i] && a[i][b[i]] == h.a[h.count+1] {
				b[i]++
				row = i
				col = b[i]
				break
			}
		}
		//某个数组遍历结束 则继续取堆顶元素
		if col >= countRow[i] {
			continue
		}
		//将数组添加到最小堆
		h.InsertData(a[row][col], MinTopSort)
	}
	t.Log(c)
	t.Log(len(c))
	t.Log(cap(c))
}
