package trietree

import "testing"

func TestBuildTree(t *testing.T) {
	root := NewTrieTree()

	BuildTree(root, "how")
	BuildTree(root, "hi")
	BuildTree(root, "her")
	BuildTree(root, "hello")
	BuildTree(root, "so")
	BuildTree(root, "se")

	pattern := "he"
	t.Log(FindTree(root, pattern))
	t.Log(FindPrefix(root, pattern))
}

func TestFirst(t *testing.T) {
	list := make([]int, 1)
	list[0] = 1
	t.Log(list)
	changeList(list)
	t.Log(list)
}

func changeList(list []int) {
	list[0] = 2
}
