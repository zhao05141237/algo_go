package trietree

import "testing"

func TestAcMatch(t *testing.T) {
	root := NewAcTree()
	BuildAcTree(root, "abcd")
	BuildAcTree(root, "bcd")
	BuildAcTree(root, "bc")
	BuildAcTree(root, "c")
	BuildFailurePointer(root, 5)
	AcMatch(root, "abcdbcdeadeaeefbcdefadeweaaddeeabcwweffbbaaajfjdkewkwncdkwkwkwksfsjfksfsaflasjfwlejewjejllsdaaxxxbbvccddeeccdebbcdeefaa")
}
