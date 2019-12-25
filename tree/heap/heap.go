package heap

import "fmt"

const (
	MaxTopSort = "asc"
	MinTopSort = "desc"
)

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
func (h *Heap) InsertData(data int, sort string) {
	//堆满
	if h.count == h.n {
		return
	}
	h.count++
	h.a[h.count] = data

	parent := h.count / 2
	i := h.count

	op := getOpBySort(sort)

	for parent > 0 {
		if compare(h.a[parent], h.a[i], op) {
			//子节点大于父节点 则交换两个节点
			h.a[parent], h.a[i] = h.a[i], h.a[parent]
			i = parent
			parent /= 2
		} else {
			break
		}
	}
}

func (h *Heap) RemoveTopData(sort string) {
	//交换根结点和最后一个节点
	if h.count == 0 {
		return
	}
	if h.count == 1 {
		//只有一个节点 删除顶结点 此处是假删除
		//h.a[h.count] = 0
		h.count--
		return
	}

	h.a[1], h.a[h.count] = h.a[h.count], h.a[1]
	h.count--
	//从顶至下最大化
	heapingUptoDown(h.a, h.count, sort)
}

func heapingUptoDown(a []int, count int, sort string) {
	parent := 1
	leftChild := 2 * parent
	RightChild := leftChild + 1
	op := getOpBySort(sort)

	for RightChild <= count && (compare(a[parent], a[leftChild], op) || compare(a[parent], a[RightChild], op)) {
		if compare(a[parent], a[leftChild], op) {
			a[parent], a[leftChild] = a[leftChild], a[parent]
		}
		if compare(a[parent], a[RightChild], op) {
			a[parent], a[RightChild] = a[RightChild], a[parent]
		}
		parent = leftChild
		leftChild = 2 * parent
		RightChild = leftChild + 1
	}
}

func compare(left int, right int, op string) bool {
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

func getOpBySort(sort string) string {
	op := "<"
	if sort == MinTopSort {
		op = ">"
	}
	return op
}

func (h *Heap) printHeap() {
	fmt.Println(h.a)
}
