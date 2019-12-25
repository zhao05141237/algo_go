package heap

import "fmt"

type Heap struct {
	a     []int
	n     int
	count int
}

//创建一个空堆
func NewHeap(n int) *Heap {
	h := &Heap{
		a:     make([]int, n+1),
		n:     n,
		count: 0,
	}
	return h
}

//插入一个元素
func (h *Heap) InsertData(data int) {
	//堆满
	if h.count == h.n {
		return
	}
	h.count++
	h.a[h.count] = data

	parent := h.count / 2
	i := h.count

	for parent > 0 {
		if h.a[parent] < h.a[i] {
			//子节点大于父节点 则交换两个节点
			h.a[parent], h.a[i] = h.a[i], h.a[parent]
			i = parent
			parent /= 2
		} else {
			break
		}
	}
}

func (h *Heap) RemoveMaxData() {
	//交换根结点和最后一个节点
	if h.count == 0 {
		return
	}
	if h.count == 1 {
		//只有一个节点 删除顶结点
		h.a[h.count] = 0
		h.count--
		return
	}

	h.a[1], h.a[h.count] = h.a[h.count], h.a[1]
	h.count--
	//从顶至下最大化
	heapingUptoDown(h.a, h.count)
}

func heapingUptoDown(a []int, count int) {
	parent := 1
	leftChild := 2 * parent
	RightChild := leftChild + 1

	for RightChild <= count && (a[parent] < a[leftChild] || a[parent] < a[RightChild]) {
		if a[parent] < a[leftChild] {
			a[parent], a[leftChild] = a[leftChild], a[parent]
		}
		if a[parent] < a[RightChild] {
			a[parent], a[RightChild] = a[RightChild], a[parent]
		}
		parent = leftChild
		leftChild = 2 * parent
		RightChild = leftChild + 1
	}
}

func (h *Heap) printHeap() {
	fmt.Println(h.a)
}
