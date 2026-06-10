# Binary Search Tree

## Validate Binary Search Tree
tree, dfs, binary-search-tree 
https://leetcode.com/problems/validate-binary-search-tree/description/

### Solution 1
Time:  O(n)
Space: O(h)
```
n = number of nodes
h = height of tree
```


```python
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:
    def isValidBST(self, root):
        def validate(node, low, high):
            if not node:
                return True

            if not (low < node.val < high):
                return False

            return (
                validate(node.left, low, node.val)
                and validate(node.right, node.val, high)
            )

        return validate(root, float("-inf"), float("inf"))
```


## Binray Tree Inorder Traversal
dfs, tree 
https://leetcode.com/problems/binary-tree-inorder-traversal/description/

### Solution 1: Recursive solution
```python
from typing import Optional, List

# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:
    def inorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        result = []

        def inorder(node):
            if not node:
                return

            inorder(node.left)
            result.append(node.val)
            inorder(node.right)

        inorder(root)
        return result
```

### Solution 2: Iterative solution
Time:  O(n)
Space: O(h)
Worst case: O(n)
Balanced tree: O(log n)

```python
from typing import Optional, List

class Solution:
    def inorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        result = []
        stack = []
        current = root

        while current or stack:
            while current:
                stack.append(current)
                current = current.left

            current = stack.pop()
            result.append(current.val)

            current = current.right

        return result
```


## Binary Search Tree Iterator
tree, binary-search-tree 
https://leetcode.com/problems/binary-search-tree-iterator/description/

### Solution 1:
For hasNext():
Time:  O(1)
Space: O(1)

For next()
Average time: O(1)
Worst-case time: O(h)

```python
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class BSTIterator:

    def __init__(self, root: Optional[TreeNode]):
        self.stack = []
        self._push_left(root)

    def _push_left(self, node):
        while node:
            self.stack.append(node)
            node = node.left

    def next(self) -> int:
        node = self.stack.pop()

        if node.right:
            self._push_left(node.right)

        return node.val

    def hasNext(self) -> bool:
        return len(self.stack) > 0
```

HasNext(): O(1)
Next(): average O(1), worst-case O(h)
Space: O(h)
```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	iterator := BSTIterator{
		stack: []*TreeNode{},
	}

	iterator.pushLeft(root)

	return iterator
}

func (this *BSTIterator) pushLeft(node *TreeNode) {
	for node != nil {
		this.stack = append(this.stack, node)
		node = node.Left
	}
}

func (this *BSTIterator) Next() int {
	// The top of the stack is the next smallest node.
	n := len(this.stack)
	node := this.stack[n-1]
	this.stack = this.stack[:n-1]

	// If this node has a right subtree,
	// the next values come from the leftmost path of that subtree.
	if node.Right != nil {
		this.pushLeft(node.Right)
	}

	return node.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}
```

## Symmetric Tree
bfs
https://leetcode.com/problems/symmetric-tree/description/

### Solution 1: Recursive solution
```python
from typing import Optional

# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:
    def isSymmetric(self, root: Optional[TreeNode]) -> bool:
        def is_mirror(left, right):
            if not left and not right:
                return True

            if not left or not right:
                return False

            if left.val != right.val:
                return False

            return (
                is_mirror(left.left, right.right)
                and is_mirror(left.right, right.left)
            )

        return is_mirror(root.left, root.right)
```

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func isSymmetric(root *TreeNode) bool {
	return isMirror(root.Left, root.Right)
}

func isMirror(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	if left.Val != right.Val {
		return false
	}

	return isMirror(left.Left, right.Right) &&
		isMirror(left.Right, right.Left)
}
```

### Solution 2: Iterative solution
```python
from collections import deque
from typing import Optional

class Solution:
    def isSymmetric(self, root: Optional[TreeNode]) -> bool:
        queue = deque()
        queue.append((root.left, root.right))

        while queue:
            left, right = queue.popleft()

            if not left and not right:
                continue

            if not left or not right:
                return False

            if left.val != right.val:
                return False

            queue.append((left.left, right.right))
            queue.append((left.right, right.left))

        return True
```

```go
func isSymmetric(root *TreeNode) bool {
	queue := []*TreeNode{root.Left, root.Right}

	for len(queue) > 0 {
		left := queue[0]
		right := queue[1]
		queue = queue[2:]

		if left == nil && right == nil {
			continue
		}

		if left == nil || right == nil {
			return false
		}

		if left.Val != right.Val {
			return false
		}

		queue = append(queue, left.Left, right.Right)
		queue = append(queue, left.Right, right.Left)
	}

	return true
}
```

