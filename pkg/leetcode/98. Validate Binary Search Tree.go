package leetcode

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var allValue []int
	allValue = isValidBSTBase(root, allValue)
	valid := true
	precede := allValue[0]
	for index, value := range allValue {
		if index == 0 {
			continue
		}
		if value <= precede {
			valid = false
			break
		}
		precede = value
	}
	return valid
}

func isValidBSTBase(root *TreeNode, allValue []int) []int {
	if root == nil {
		return allValue
	}
	if root.Left != nil {
		allValue = isValidBSTBase(root.Left, allValue)
	}
	allValue = append(allValue, root.Val)
	if root.Right != nil {
		allValue = isValidBSTBase(root.Right, allValue)
	}
	return allValue
}

/**
class Solution {
public:
    TreeNode* pre=nullptr;
    bool isValidBST(TreeNode* root) {
        if (!root) return true;
        if (!isValidBST(root->left)) return false;
        if (pre && root->val <= pre->val) return false;
        pre=root;
        if (!isValidBST(root->right)) return false;
        return true;
    }
};
使用一个pre来作为前一个数字的指针
*/

/**
if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}
	valid := true
	if root.Left != nil {
		valid = valid && root.Val > root.Left.Val
		if valid == false {
			return valid
		}
		valid = valid && isValidBSTBase(root.Left)
	}
	if root.Right != nil {
		valid = valid && root.Val < root.Right.Val
		if valid == false {
			return valid
		}
		valid = valid && isValidBSTBase(root.Right)
	}
	return valid
*/

type Stack []interface{}

func (s *Stack) Push(value interface{}) {
	*s = append(*s, value)
}

func (s *Stack) Pop() (interface{}, bool) {
	stack := *s
	if len(stack) == 0 {
		return nil, false // 栈为空
	}
	value := stack[len(stack)-1]
	*s = stack[:len(stack)-1]
	return value, true
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}
