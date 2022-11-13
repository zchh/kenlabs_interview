package s3

type Tree struct {
	Val   string
	Left  *Tree
	Right *Tree
}

func S3(t1 *Tree, t2 *Tree) *Tree {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.Val += t2.Val
	t1.Left = S3(t1.Left, t2.Left)
	t1.Right = S3(t1.Right, t2.Right)
	return t1
}
