package s1

type Tree struct {
	Val   string
	Left  *Tree
	Right *Tree
}

func S1(root1 *Tree, root2 *Tree) bool {
	if root1 == root2 {
		return true
	}
	if root1 == nil || root2 == nil || root1.Val != root2.Val {
		return false
	}
	return (S1(root1.Left, root2.Left) && S1(root1.Right, root2.Right) ||
		S1(root1.Left, root2.Right) && S1(root1.Right, root2.Left))
}
