package medianFinder

import (
	"fmt"
)

/**
中位数是有序列表中间的数。如果列表长度是偶数，中位数则是中间两个数的平均值。
例如，
[2,3,4] 的中位数是 3
[2,3] 的中位数是 (2 + 3) / 2 = 2.5
设计一个支持以下两种操作的数据结构：
void addNum(int num) - 从数据流中添加一个整数到数据结构中。
double findMedian() - 返回目前所有元素的中位数。
示例：
addNum(1)
addNum(2)
findMedian() -> 1.5
addNum(3)
findMedian() -> 2
进阶:
如果数据流中所有整数都在 0 到 100 范围内，你将如何优化你的算法？
如果数据流中 99% 的整数都在 0 到 100 范围内，你将如何优化你的算法？
*/

const (
	MaxTopSort = "asc"
	MinTopSort = "desc"
)

type MedianFinder struct {
	count int
	max   *Heap
	min   *Heap
}

type Heap struct {
	a     []int
	count int
	sort  string
}

type Hearer interface {
	NewHeap() *Heap
}

type MaxHeap struct {
}

type MinHeap struct {
}

func (m *MaxHeap) NewHeap() *Heap {
	return &Heap{
		a:     make([]int, 1),
		count: 0,
		sort:  MaxTopSort,
	}
}

func (m *MinHeap) NewHeap() *Heap {
	return &Heap{
		a:     make([]int, 1),
		count: 0,
		sort:  MinTopSort,
	}
}

func insertData(h *Heap, data int) {
	h.count++
	h.a = append(h.a, data)

	parent := h.count / 2
	i := h.count

	op := GetOpBySort(h.sort)

	for parent > 0 {
		if Compare(h.a[parent], h.a[i], op) {
			//子节点大于父节点 则交换两个节点
			h.a[parent], h.a[i] = h.a[i], h.a[parent]
			i = parent
			parent /= 2
		} else {
			break
		}
	}
}

func Compare(left int, right int, op string) bool {
	switch op {
	case ">":
		return left > right
	case "<":
		return left < right
	case ">=":
		return left >= right
	case "<=":
		return left <= right
	default:
		return false
	}
}

func GetOpBySort(sort string) string {
	op := "<"
	if sort == MinTopSort {
		op = ">"
	}
	return op
}

func HeapingUptoDown(a []int, count int, sort string) {
	op := GetOpBySort(sort)
	for parent := 1; parent <= count/2; {
		index := parent
		if Compare(a[parent], a[parent*2], op) {
			index = parent * 2
		}
		if 2*parent+1 <= count && Compare(a[index], a[parent*2+1], op) {
			index = parent*2 + 1
		}
		if index == parent {
			break
		}
		a[parent], a[index] = a[index], a[parent]
		parent = index
	}
}

func removeTopData(h *Heap) {
	//交换根结点和最后一个节点
	if h.count == 0 {
		return
	}
	if h.count == 1 {
		//只有一个节点 删除顶结点 此处是假删除
		h.count--
		return
	}

	h.a[1], h.a[h.count] = h.a[h.count], h.a[1]
	h.count--
	h.a = h.a[0 : h.count+1]
	//从顶至下最大化
	HeapingUptoDown(h.a, h.count, h.sort)
}

func printHeap(h *Heap) {
	fmt.Println(h.a)
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		count: 0,
		max:   (&MaxHeap{}).NewHeap(),
		min:   (&MinHeap{}).NewHeap(),
	}
}

func (m *MedianFinder) AddNum(num int) {
	m.count++
	if m.count == 1 {
		//初始情况加到大顶堆
		insertData(m.max, num)
	} else {
		if m.max.a[1] < num {
			//如果插入数据大于大顶堆堆顶元素 则插入最小堆
			insertData(m.min, num)
		} else {
			//如果插入数据小于等于大顶堆堆顶元素 则插入最大堆
			insertData(m.max, num)
		}
		//此时需要调整两个堆 满足总数是偶数时大顶堆和小顶堆个数均为n/2 奇数时大顶堆为n/2 + 1 小顶堆为n/2 不满足则移除堆顶数
		if m.count%2 == 0 {
			if m.max.count > m.count/2 {
				insertData(m.min, m.max.a[1])
				removeTopData(m.max)
			} else if m.max.count < m.count/2 {
				insertData(m.max, m.min.a[1])
				removeTopData(m.min)
			}
		} else {
			if m.max.count > m.count/2+1 {
				insertData(m.min, m.max.a[1])
				removeTopData(m.max)
			} else if m.max.count < m.count/2+1 {
				insertData(m.max, m.min.a[1])
				removeTopData(m.min)
			}
		}
	}

}

func (m *MedianFinder) FindMedian() float64 {
	if m.count == 0 {
		return 0
	} else if m.count == 1 {
		return float64(m.max.a[1])
	} else if m.count%2 == 0 {
		return float64(m.max.a[1]+m.min.a[1]) / float64(2)
	} else {
		return float64(m.max.a[1])
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
