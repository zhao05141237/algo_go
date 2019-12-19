package recursiveReferer

/**
给定一个用户 ID，如何查找这个用户的“最终推荐人”
*/
import (
	"testing"
)

func findReferer(a map[string]string, b string) string {
	if val, ok := a[b]; ok == false {
		return b
	} else {
		return findReferer(a, val)
	}
}

func TestFirst(t *testing.T) {
	a := map[string]string{"B": "A", "C": "B"}
	val := findReferer(a, "C")
	t.Log(val)
}

func TestSec(t *testing.T) {
	a := map[string]string{"B": "A", "C": "B", "D": "C", "F": "D"}
	val := findReferer(a, "F")
	t.Log(val)
}
