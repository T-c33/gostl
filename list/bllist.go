package list

type BNode struct {
	value interface{}
	next *BNode
	prev *BNode
}

type BList struct {
	root *BNode
	length int
}
