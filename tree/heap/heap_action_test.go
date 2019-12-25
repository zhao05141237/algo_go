package heap

import (
	"math/rand"
	"testing"
	"time"
)

func TestHeap_InsertData(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	h := NewHeap(10)
	for i := 1; i <= 10; i++ {
		h.InsertData(rand.Intn(100))
	}
	//a := []int{82,79,77,24,25,3,51,14,12,2}
	//for _, value := range a {
	//	h.InsertData(value)
	//}
	h.printHeap()
	for i := 1; i < 10; i++ {
		h.RemoveMaxData()
	}
	h.printHeap()

}
