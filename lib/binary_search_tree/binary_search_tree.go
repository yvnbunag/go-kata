package lib_binary_search_tree

type TreeNode struct {
	Value int
	left  *TreeNode
	right *TreeNode
}

func (treeNode *TreeNode) Insert(value int) bool {
	if treeNode.Value > value {
		if treeNode.left == nil {
			treeNode.left = &TreeNode{Value: value}
			return true
		}

		return treeNode.left.Insert(value)
	}

	if treeNode.Value < value {
		if treeNode.right == nil {
			treeNode.right = &TreeNode{Value: value}
			return true
		}
		return treeNode.right.Insert(value)
	}

	return false
}

func (treeNode *TreeNode) ToList() (result []int) {
	if treeNode.left != nil {
		result = append(result, treeNode.left.ToList()...)
	}

	result = append(result, treeNode.Value)

	if treeNode.right != nil {
		result = append(result, treeNode.right.ToList()...)
	}

	return
}

func (treeNode *TreeNode) Get(value int) *TreeNode {
	if treeNode.Value == value {
		return treeNode
	}

	if treeNode.Value > value && treeNode.left != nil {
		return treeNode.left.Get(value)
	}

	if treeNode.Value < value && treeNode.right != nil {
		return treeNode.right.Get(value)
	}

	return nil
}
