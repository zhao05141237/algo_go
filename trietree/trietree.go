package trietree

type TrieNode struct {
	data         byte
	children     []*TrieNode
	isEndingChar bool
}

func NewTrieTree() *TrieNode {
	t := &TrieNode{
		data:         '/',
		children:     make([]*TrieNode, 26),
		isEndingChar: false,
	}
	return t
}

func BuildTree(t *TrieNode, text string) {
	if t == nil {
		return
	}
	p := t
	textS := []byte(text)
	for _, val := range textS {
		index := val - 'a'
		if index >= 26 || index < 0 {
			return
		}
		if p.children[index] == nil {
			p.children[index] = &TrieNode{
				data:         val,
				children:     make([]*TrieNode, 26),
				isEndingChar: false,
			}
		}
		p = p.children[index]
	}
	p.isEndingChar = true
}

func FindTree(t *TrieNode, pattern string) bool {
	if t == nil {
		return false
	}

	p := t
	textS := []byte(pattern)

	for _, val := range textS {
		index := val - 'a'
		if index >= 26 || index < 0 {
			return false
		}
		if p.children[index] != nil {
			p = p.children[index]
		} else {
			return false
		}
	}

	if p.isEndingChar {
		return true
	} else {
		return false
	}

}

/**
查找pattern前缀的所有字符串
*/
func FindPrefix(t *TrieNode, pattern string) []string {
	if t == nil || len(pattern) == 0 {
		return nil
	}

	p := t
	textS := []byte(pattern)

	for _, val := range textS {
		index := val - 'a'
		if index >= 26 || index < 0 {
			return nil
		}
		if p.children[index] != nil {
			p = p.children[index]
		} else {
			return nil
		}
	}
	lists := make([]string, 0)
	length := len(pattern)
	var before string
	if length > 1 {
		before = string([]byte(pattern)[0 : length-1])
	}
	lists = printChildVal(p, before, lists)
	return lists
}

func printChildVal(t *TrieNode, before string, lists []string) []string {
	if t.isEndingChar == true {
		lists = append(lists, before+string(t.data))
		return lists
	}
	for _, value := range t.children {
		if value != nil {
			lists = printChildVal(value, before+string(t.data), lists)
		}
	}
	return lists
}
