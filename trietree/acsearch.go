package trietree

import (
	"algo/queue/queueArray"
	"fmt"
)

type AcNode struct {
	data         byte
	children     []*AcNode
	isEndingChar bool
	length       int
	fail         *AcNode
}

func NewAcTree() *AcNode {
	t := &AcNode{
		data:         '/',
		children:     make([]*AcNode, 26),
		isEndingChar: false,
		length:       0,
		fail:         nil,
	}
	return t
}

func CreateAcNode(data byte) *AcNode {
	return &AcNode{
		data:         data,
		children:     make([]*AcNode, 26),
		isEndingChar: false,
		length:       0,
		fail:         nil,
	}
}

//构建模式串tree
func BuildAcTree(t *AcNode, text string) {
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
			p.children[index] = CreateAcNode(val)
			p.children[index].length = p.length + 1
		}
		p = p.children[index]
	}
	p.isEndingChar = true
}

func BuildFailurePointer(t *AcNode, n int) {
	queue := queueArray.NewArrayQueue(n)
	queue.EnQueue(t)
	for !queue.IsEmpty() {
		p := queue.DeQueue()
		for i := 0; i < 26; i++ {
			pc := p.(*AcNode).children[i]
			if pc == nil {
				continue
			}
			if p == t {
				pc.fail = t
			} else {
				q := p.(*AcNode).fail
				for q != nil {
					qc := q.children[pc.data-'a']
					if qc != nil {
						pc.fail = qc
						break
					}
					q = q.fail
				}
				if q == nil {
					pc.fail = t
				}
			}
			queue.EnQueue(pc)
		}
	}
}

func AcMatch(t *AcNode, text string) {
	fmt.Printf("原始主串%s\n:", text)
	textLength := len(text)
	textS := []byte(text)
	p := t
	for i := 0; i < textLength; i++ {
		index := textS[i] - 'a'
		for p.children[index] == nil && p != t {
			p = p.fail
		}
		p = p.children[index]
		if p == nil {
			p = t
		}
		tmp := p
		for tmp != t {
			if tmp.isEndingChar {
				pos := i - tmp.length + 1
				for j := pos; j < pos+tmp.length; j++ {
					//敏感词用*号代替
					textS[j] = '*'
				}
				fmt.Printf("匹配起始下标%d;长度%d\n", pos, tmp.length)
			}
			tmp = tmp.fail
		}
	}
	fmt.Printf("替换后主串:%s\n", textS)
}
